package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/UmetsuJunya/todo-app-backend/backend/todo"

	"github.com/gin-gonic/gin"
)

func TodosGet(todoAPI *todo.Todo) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := todoAPI.GetAll()
		c.JSON(http.StatusOK, result)
	}
}

type TodoPostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

func TodoPost(todoAPI *todo.Todo) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := TodoPostRequest{}

		if err := c.Bind(&requestBody); err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
			return
		}

		item := todo.Todo{
			Title:       requestBody.Title,
			Description: requestBody.Description,
			// Status:      requestBody.Status,
		}
		result := todoAPI.Add(item)
		sendBackResult(c, result)
	}
}

type TodoPutRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

func TodoPut(todoAPI *todo.Todo) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := TodoPutRequest{}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		if err := c.Bind(&requestBody); err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
			return
		}

		item := todo.Todo{
			Title:       requestBody.Title,
			Description: requestBody.Description,
			// Status:      requestBody.Status,
		}
		result := todoAPI.UpdateTodo(id, item)
		sendBackResult(c, result)
	}
}

func TodoComplete(todoAPI *todo.Todo) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("pdpdpdpdppd")
		fmt.Println(*todoAPI)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}
		result := todoAPI.UpdateStatus(id)
		sendBackResult(c, result)
	}
}

func TodoDelete(todoAPI *todo.Todo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		result := todoAPI.Delete(id)
		sendBackResult(c, result)
	}
}

func sendBackResult(c *gin.Context, result bool) {
	if result {
		c.JSON(http.StatusOK, gin.H{"message": "Operation succeeded"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Operation failed"})
	}
}
