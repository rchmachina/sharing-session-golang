package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rchmachina/sharing-session-golang/model"
	"github.com/rchmachina/sharing-session-golang/repositories"
	routes "github.com/rchmachina/sharing-session-golang/route"
	"github.com/rchmachina/sharing-session-golang/utils/database"
)

func main() {
	// Load the environment variables from the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	// Initialize the Gin router
	r := gin.Default()

	// CORS middleware for API
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	// Connect to the database
	db := database.DatabaseConnection()
	fmt.Println("Connected to database")

	// Test model data
	modelTesting := model.CreateUser{
		UserName:       "testing json1",
		HashedPassword: "dwaarrr",
		Roles: "mimin2",
	}

	// Test repository
	testRepositoryUser := repositories.RepositoryUser(db)
	testRepositoryUser.CreateUserJson(modelTesting)

	// Initialize routes
	routes.RouteInit(r.Group("/api/V1"))

	// Start the server
	port := os.Getenv("PORT")
	fmt.Println("Server running on localhost:", port)
	if err := r.Run( port); err != nil {
		log.Fatal(err)
	}
}
