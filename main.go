package main

import (
	"gocrudapibackend/controllers"
	"gocrudapibackend/initializers"
	"gocrudapibackend/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.EnvVariable()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run() // listen and serve on 0.0.0.0:8080
}
