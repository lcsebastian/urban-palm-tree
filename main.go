package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lcsebastian/urban-palm-tree/cmd"
	scraper "github.com/lcsebastian/urban-palm-tree/cmd/scraper"
)

func getResumes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cmd.LoadTestResumes("data/resume.json"))
}

func getFilters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cmd.LoadTestFilters("data/filter.json"))
}

func getJobs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, scraper.CollyScrape())
}

func main() {

	fmt.Println("Setting up router")
	router := gin.Default()
	router.GET("/resumes", getResumes)
	router.GET("/filters", getFilters)
	router.GET("/jobs", getJobs)
	router.Run("localhost:8080")
}
