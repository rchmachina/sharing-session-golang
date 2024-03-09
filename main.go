package main

import (

	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	routes "github.com/rchmachina/sharing-session-golang/route"
	"github.com/rchmachina/sharing-session-golang/utils/database"
	flags "github.com/rchmachina/sharing-session-golang/utils/flag"
	"log"
	"net/http"
	"os"
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

	// check connection to  database
	_ = database.DatabaseConnection()
	fmt.Println("Connected to database")
	flags.MigrationFlags()



	// Initialize routes
	APIVersion := os.Getenv("API_VERSION")
	routes.RouteInit(r.Group(APIVersion))

	// Start the server
	port := os.Getenv("PORT")
	fmt.Println("Server running on localhost:", port)
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
