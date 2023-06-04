package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kundu-ramit/dozer_match/domain/entity"
	"github.com/kundu-ramit/dozer_match/service"
)

type ScraperController interface {
	Get(c *gin.Context)
	StartScrape(c *gin.Context)
	Clear(c *gin.Context)
}

type scraperController struct {
	service service.DozerService
}

func NewScraperController() ScraperController {
	return &scraperController{
		service: service.NewDozerService(),
	}
}

func (sc *scraperController) Get(c *gin.Context) {
	scrapeId := c.Param("id")

	var dozers []entity.BullDozer
	var err error
	if scrapeId == "" {
		dozers, err = sc.service.FetchLatest(c)
	} else {
		dozers, err = sc.service.FetchById(c, scrapeId)
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, dozers)
}

func (sc *scraperController) StartScrape(c *gin.Context) {
	scrapeIndex := sc.generateScrapeIndex()

	err := sc.service.CheckExistingScrape(c)
	if err != nil {
		dozers, err := sc.service.FetchLatest(c)
		if err == nil {
			c.JSON(http.StatusOK, dozers)
			return
		}
	}

	dozers, err := sc.service.StartScrape(c, scrapeIndex)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, dozers)
}

func (sc *scraperController) Clear(c *gin.Context) {

	err := sc.service.Delete(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, `All scrapes have been deleted`)
}

func (sc *scraperController) generateScrapeIndex() string {
	rand.Seed(time.Now().UnixNano())
	min := 1000000000
	max := 9999999999
	scrapeIndex := rand.Intn(max-min+1) + min

	return fmt.Sprintf("%010d", scrapeIndex)
}
