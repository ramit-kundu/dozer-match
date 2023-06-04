package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	controllers "github.com/kundu-ramit/dozer_match/controller"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	scraperController := controllers.NewScraperController()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "healthy",
		})
	})

	//find a specific scrape
	router.GET("/scrape/:id", scraperController.Get)

	//find the latest scrape
	router.GET("/scrape", scraperController.Get)

	//order a new scrape
	router.POST("/scrape", scraperController.StartScrape)

	//delete all scrape records
	router.DELETE("/scrape", scraperController.Clear)

	return router
}
