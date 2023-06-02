package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kundu-ramit/dozer_match/service"
)

type ScraperController struct {
	service service.DozerService
}

func (sc *ScraperController) Get(c *gin.Context) {

	c.JSON(http.StatusOK, scrapeResults)
}

func (sc *ScraperController) StartScrape(c *gin.Context) {

	resp, err := http.Get(serviceURL)
	if err != nil {
		// Handle the error appropriately
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start scrape"})
		return
	}

	var response struct {
		Success bool `json:"success"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		// Handle the error appropriately
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse service response"})
		return
	}

	if !response.Success {
		// Handle the unsuccessful response from the service
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Scrape service returned an error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Scraping started successfully"})
}
