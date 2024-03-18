package user

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"greatcomcatengineering.com/backend/database"
	"greatcomcatengineering.com/backend/models"
	"greatcomcatengineering.com/backend/utils"
	"net/http"
	"strings"
)

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	ctx := context.TODO()

	// Generate a UUID for the new user
	newUser.ID = uuid.New().String()

	// Attempt to add the new user to the database
	if err := db.AddUser(ctx, newUser); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create user: "+err.Error())
		return
	}

	// Successfully created the user, respond with the new user object
	utils.RespondWithJSON(w, http.StatusCreated, "User created successfully", newUser)
}

func HandleGetUserByEmail(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL path

	pathParts := strings.Split(r.URL.Path, "/")
	var id string
	if len(pathParts) > 0 {
		id = pathParts[len(pathParts)-1] // Get the ID part
	}

	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	ctx := context.TODO()

	// Attempt to retrieve the user from the database
	user, err := db.GetUserByEmail(ctx, id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve user: "+err.Error())
		return
	}

	// Successfully retrieved the user, respond with the user object
	utils.RespondWithJSON(w, http.StatusOK, "User retrieved successfully", user)
}
