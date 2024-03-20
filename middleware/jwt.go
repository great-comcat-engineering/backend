package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"greatcomcatengineering.com/backend/configs"
	"greatcomcatengineering.com/backend/models"
	"greatcomcatengineering.com/backend/utils"
)

var jwtSecret = []byte(configs.AppConfig().Auth.JWTSecret)

func GenerateJWTToken(email string, accountType models.UserType) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Create a map to store our claims
	claims := token.Claims.(jwt.MapClaims)

	// Set token claims
	claims["iss"] = "greatcomcatengineering.com"
	claims["accountType"] = accountType
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(configs.AppConfig().Auth.JWTExpireInHours)).Unix() // Token expires after 24 hours

	// Generate encoded token and send it as response.
	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, err
}

// JWTAuthMiddleware checks for a valid JWT token in the Authorization header.
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")

		if !strings.HasPrefix(authHeader, BearerSchema) {
			utils.RespondWithError(c, http.StatusUnauthorized, "Authorization header must start with Bearer")
			return
		}

		tokenString := authHeader[len(BearerSchema):]
		fmt.Println("->" + tokenString)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		fmt.Println(token)
		fmt.Println(err)

		if err != nil {
			var errMsg string
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					errMsg = "Malformed token"
				} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
					errMsg = "Token is expired"
				} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
					errMsg = "Token not valid yet"
				} else {
					errMsg = "Invalid token"
				}
			} else {
				errMsg = "Invalid token"
			}
			utils.RespondWithError(c, http.StatusUnauthorized, errMsg)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("email", claims["email"])
			c.Set("accountType", claims["accountType"])
		} else {
			utils.RespondWithError(c, http.StatusUnauthorized, "Invalid token")
			return
		}

		c.Next()
	}
}

// Check if the user is an admin
func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		accountType := c.GetString("accountType")
		if accountType != string(models.AdminUser) {
			utils.RespondWithError(c, http.StatusUnauthorized, "Unauthorized")
			return
		}
		c.Next()
	}
}
