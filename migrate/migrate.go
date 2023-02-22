package main

import (
	"gocrudapibackend/initializers"
	"gocrudapibackend/models"
)

func init() {
	initializers.EnvVariable()
	initializers.ConnectToDb()
}

func main() {

	// Migrate the schema
	initializers.DB.AutoMigrate(&models.User{})

}
