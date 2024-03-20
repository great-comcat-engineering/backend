package utils

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"greatcomcatengineering.com/backend/models"
	"io"
	"log"
)

// RespondWithJSON is a utility function to respond with JSON to the client with the given status code, message, and content
func RespondWithJSON[T any](c *gin.Context, status int, message string, content T) {
	response := models.ApiResponse[T]{
		Status:  status,
		Message: message,
		Data:    content,
	}
	c.JSON(status, response)
}

// RespondWithError is a utility function to respond with an error message to the client with the given status code and message
func RespondWithError(c *gin.Context, status int, message string) {
	RespondWithJSON[interface{}](c, status, message, nil)
	c.Abort()
}

// LogRequestBody is a middleware function to log the request body before passing it to the next middleware/handler
func LogRequestBody(c *gin.Context) {
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = io.ReadAll(c.Request.Body)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	var bodyMap map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &bodyMap); err == nil {
		bodyString, _ := json.Marshal(bodyMap)
		log.Println("Request body --> ", string(bodyString))
	} else {
		log.Println("Error reading request body for logging")
	}

	c.Next()
}

func GenerateID() string {
	return uuid.New().String()
}
