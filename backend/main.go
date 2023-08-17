package main

import (
	"os"
	"time"

	"github.com/UmetsuJunya/todo-app-backend/backend/handler"
	"github.com/UmetsuJunya/todo-app-backend/backend/lib"
	"github.com/UmetsuJunya/todo-app-backend/backend/todo"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/gin-contrib/cors"
)

func main() {
	if os.Getenv("USE_HEROKU") != "1" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	todo := todo.New()
	// user := user.New()

	lib.DBOpen()
	defer lib.DBClose()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	r.GET("/todo", handler.TodosGet(todo))
	r.POST("/todo", handler.TodoPost(todo))
	r.PUT("/todo/update/:id", handler.TodoPut(todo))
	r.PUT("/todo/switch/:id", handler.TodoComplete(todo))
	r.DELETE("/todo/delete/:id", handler.TodoDelete(todo))
	// r.POST("/user/login", handler.UserPost(user))

	r.Run(os.Getenv("HTTP_HOST") + ":" + os.Getenv("HTTP_PORT")) // listen and serve on 0.0.0.0:8080
}
