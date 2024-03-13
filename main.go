package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{
		ID:        "1",
		Item:      "Learn Go 1",
		Completed: false,
	},
	{
		ID:        "2",
		Item:      "Learn Go 2",
		Completed: false,
	},
	{
		ID:        "3",
		Item:      "Learn Go 3",
		Completed: false,
	},
}

// getTodos returns the list of todos
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

// addTodo adds a new todo to the list
func addTodo(context *gin.Context) {

	fmt.Println("addTodo")
	fmt.Println(context)

	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		todos = append(todos, newTodo)
		context.IndentedJSON(http.StatusCreated, newTodo)
	}
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)

	router.POST("/todos", addTodo)

	router.Run("localhost:9090")
}
