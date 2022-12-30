package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string
	Item      string
	Completed bool
}

var todos = []todo{
	{ID: "1", Item: "Study", Completed: false},
	{ID: "2", Item: "Workout", Completed: false},
	{ID: "3", Item: "Work", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodos(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)

}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodosById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)

}

func getTodosById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("item Not found")
}

func updateTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodosById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", updateTodoStatus)
	router.POST("/todos", addTodos)
	router.Run("localhost:8000")
}
