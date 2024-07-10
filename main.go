package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getResumes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, loadTestResumes("data/resume.json"))
}

func getFilters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, loadTestFilters("data/filter.json"))
}

func main() {
	fmt.Println("Setting up router")
	router := gin.Default()
	router.GET("/resumes", getResumes)
	router.GET("/filters", getFilters)
	router.Run("localhost:8080")
}
