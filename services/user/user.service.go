package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"greatcomcatengineering.com/backend/database"
	"greatcomcatengineering.com/backend/middleware"
	"greatcomcatengineering.com/backend/models"
	"greatcomcatengineering.com/backend/utils"
	"log"
	"net/http"
)

// @Summary Create new user
// @Description Creates a new user with the provided information.
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user  body      models.RegisterRequest   true  "User Data"
// @Success 201  {object}  models.User  "User created successfully"
// @Failure 400  {object}  nil  "Invalid request payload"
// @Failure 500  {object}  nil  "Internal Server Error"
// @Router /user/create [post]
func HandleCreateUser(c *gin.Context) {
	var req models.RegisterRequest

	utils.LogRequestBody(c)

	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if errMsg, valid := ValidateUserRegister(req); !valid {
		utils.RespondWithError(c, http.StatusBadRequest, errMsg)
		return
	}

	doesExist, _ := database.UserExists(c.Request.Context(), req.Email)
	if doesExist {
		utils.RespondWithError(c, http.StatusBadRequest, "User with this email already exists")
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error hashing password")
		return
	}

	user := models.User{
		Email:       req.Email,
		Password:    string(hashed),
		AccountType: models.DefaultUser,
	}

	if err := database.AddUser(c.Request.Context(), user); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.RespondWithJSON(c, http.StatusCreated, "User created successfully", user)
}

// @Summary Get user by email
// @Description Retrieves a user by the provided email.
// @Tags users
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param   email  path  string  true  "User Email"
// @Success 200  {object}  models.User  "User retrieved successfully"
// @Failure 400  {object}  nil  "Invalid or missing email"
// @Failure 404  {object}  nil  "User not found"
// @Failure 500  {object}  nil  "Internal Server Error"
// @Router /user/{email} [get]
func HandleGetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid or missing email")
		return
	}

	user, err := database.GetUserByEmail(c.Request.Context(), email)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to retrieve user")
	} else if user == nil {
		utils.RespondWithError(c, http.StatusNotFound, "User not found")
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, "User retrieved successfully", user)
}

// @Summary Get all users
// @Description Retrieves all users from the database.
// @Tags users
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200  {array}  models.User  "Users retrieved successfully"
// @Failure 401  {object}  nil  "Unauthorized access"
// @Failure 500  {object}  nil  "Internal Server Error"
// @Router /user/all [get]
func HandleGetAllUsers(c *gin.Context) {
	users, err := database.GetAllUsers(c.Request.Context())
	if err != nil {
		// Log the error for internal tracking without exposing sensitive details to the client
		log.Printf("Error retrieving all users: %v", err)
		utils.RespondWithError(c, http.StatusInternalServerError, "An error occurred while processing your request")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, "Users retrieved successfully", users)
}

// @Summary Login
// @Description Logs in the user with the provided email and password.
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user  body  models.LoginRequest  true  "User Data"
// @Success 200  {object}  map[string]interface{} "Login successful, token returned"
// @Failure 400  {object}  nil  "Invalid request payload"
// @Failure 401  {object}  nil  "Invalid email or password"
// @Failure 500  {object}  nil  "Internal Server Error"
// @Router /user/login [post]
func HandleLogin(c *gin.Context) {
	var loginRequest models.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	dbUser, err := database.GetUserByEmail(c.Request.Context(), loginRequest.Email)
	if err != nil {
		log.Printf("Error retrieving user by email: %v, error: %v", loginRequest.Email, err)
		utils.RespondWithError(c, http.StatusInternalServerError, "An error occurred while processing your request")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginRequest.Password)); err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	token, err := middleware.GenerateJWTToken(dbUser.Email, dbUser.AccountType)
	if err != nil {
		// Log this error for internal tracking
		log.Printf("Error generating JWT token for user: %v, error: %v", dbUser.Email, err)
		utils.RespondWithError(c, http.StatusInternalServerError, "An error occurred while processing your request")
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, "Login successful", map[string]interface{}{"token": token})
}

// @Summary Get current user
// @Description Retrieves the current user from the database.
// @Tags users
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200  {object}  models.User  "User retrieved successfully"
// @Failure 401  {object}  nil  "Unauthorized or invalid user ID"
// @Failure 500  {object}  nil  "Internal Server Error"
// @Router /user/current [get]
func HandleGetCurrentUser(c *gin.Context) {
	email, exists := c.Get("email")
	if !exists {
		utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized - Unable to obtain user information from the token.")
		return
	}

	userEmail, ok := email.(string)
	if !ok {
		utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized - User email in token is invalid.")
		return
	}

	user, err := database.GetUserByEmail(c.Request.Context(), userEmail)
	if err != nil {
		log.Printf("Error retrieving user by email: %v, error: %v", userEmail, err)
		utils.RespondWithError(c, http.StatusInternalServerError, "Internal Server Error - Unable to retrieve user information.")
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, "User retrieved successfully", user)
}
