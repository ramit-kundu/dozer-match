package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	catscraper "github.com/kundu-ramit/dozer_match/domain/scraper/cat_scraper"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Define your routes here
	router.GET("/users", func(c *gin.Context) {
		// Handle GET request to /users
		c.JSON(200, gin.H{"message": "Get all users"})
		catscraper.Crawl(c)
	})

	router.POST("/users", func(c *gin.Context) {
		// Handle POST request to /users
		c.JSON(201, gin.H{"message": "Create a new user"})
	})

	// Return the router instance
	return router
}
