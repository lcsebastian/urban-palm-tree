package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/lcsebastian/urban-palm-tree/cmd"
)

func getResumes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cmd.LoadTestResumes("data/resume.json"))
}

func getFilters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cmd.LoadTestFilters("data/filter.json"))
}

func main() {
	fmt.Println("Setting up router")
	router := gin.Default()
	router.GET("/resumes", getResumes)
	router.GET("/filters", getFilters)
	router.Run("localhost:8080")
}
