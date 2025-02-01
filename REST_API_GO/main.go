package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

type todoHandlerInterface interface {
	getTodos(c *gin.Context)
	getTodoByID(c *gin.Context)
	createTodo(c *gin.Context)
	changeTodo(c *gin.Context)
	deleteTodo(c *gin.Context)
}

type todoHandler struct {
	todos []todo
	todoHandlerInterface
}

func newTodoHandler() *todoHandler {
	return &todoHandler{
		todos: []todo{
			{ID: "1", Item: "Learn Go", Completed: false},
			{ID: "2", Item: "Build REST API", Completed: false},
			{ID: "3", Item: "Learn Docker", Completed: false},
		},
	}
}

func (h *todoHandler) getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, h.todos)
}

func (h *todoHandler) getTodoByID(c *gin.Context) {
	id := c.Param("id")
	for _, todo := range h.todos {
		if todo.ID == id {
			c.IndentedJSON(http.StatusOK, todo)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

// POST
func (h *todoHandler) createTodo(c *gin.Context) {
	var newTodo todo
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	h.todos = append(h.todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

// PATCH
func (h *todoHandler) changeTodo(c *gin.Context) {
	id := c.Param("id")
	for i, todo := range h.todos {
		if todo.ID == id {
			h.todos[i].Completed = !h.todos[i].Completed
			c.IndentedJSON(http.StatusOK, h.todos[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})

}

// DELETE
func (h *todoHandler) deleteTodo(c *gin.Context) {
	id := c.Param("id")
	for i, todo := range h.todos {
		if todo.ID == id {
			// 删除元素
			h.todos = append(h.todos[:i], h.todos[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "todo deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func main() {
	r := gin.Default()
	th := newTodoHandler()

	r.GET("/todo", th.getTodos)
	r.GET("/todo/:id", th.getTodoByID)
	r.POST("/todo", th.createTodo)
	r.PATCH("/todo/:id", th.changeTodo)
	r.DELETE("/todo/:id", th.deleteTodo)

	r.Run(":8080")
}
