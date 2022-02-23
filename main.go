package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID			string	`json:"id"`
	Title		string	`json:"title"`
	Completed 	bool	`json:"completed"`
}

var todos = []todo{
	{ID: "1", Title: "In Search of Lost Time", Completed: false},
	{ID: "2", Title: "The Great Gatsby", Completed: false},
	{ID: "3", Title: "War and Peace", Completed: false},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context){
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func todoById(c *gin.Context) {
	id := c.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		c.IndentedJSON(http.StatusTeapot, gin.H{"message": "Todo not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*todo, error) {
	for i, b := range todos {
		if b.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.Run("localhost:8081")
}