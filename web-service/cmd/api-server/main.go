package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type Resume struct {
	ID        int `json:"id,omitempty"`
	Education []struct {
		Institution string `json:"institution,omitempty"`
		StudyType   string `json:"studyType,omitempty"`
		Area        string `json:"area,omitempty"`
		Score       string `json:"score,omitempty"`
		Date        string `json:"date,omitempty"`
		Summary     string `json:"summary,omitempty"`
	} `json:"education,omitempty"`
	Certifications []struct {
		Name    string `json:"name,omitempty"`
		Issuer  string `json:"issuer,omitempty"`
		Date    string `json:"date,omitempty"`
		Summary string `json:"summary,omitempty"`
	} `json:"certifications,omitempty"`
	Experience []struct {
		Company  string `json:"company,omitempty"`
		Position string `json:"position,omitempty"`
		Location string `json:"location,omitempty"`
		Date     string `json:"date,omitempty"`
		Summary  string `json:"summary,omitempty"`
	} `json:"experience,omitempty"`
	Projects []struct {
		Name        string `json:"name,omitempty"`
		Date        string `json:"date,omitempty"`
		Description string `json:"description,omitempty"`
		Website     string `json:"website,omitempty"`
		Summary     string `json:"summary,omitempty"`
	} `json:"projects,omitempty"`
	Skills []struct {
		Name  string `json:"name,omitempty"`
		Level string `json:"level,omitempty"`
	} `json:"skills,omitempty"`
}

type Filter struct {
	ID         int      `json:"id,omitempty"`
	DatePosted string   `json:"date posted,omitempty"`
	RemoteOnly bool     `json:"remote only,omitempty"`
	Location   []string `json:"location,omitempty"`
	JobType    []string `json:"job type,omitempty"`
	MinimumPay int      `json:"minimum pay,omitempty"`
	Keywords   []string `json:"keywords,omitempty"`
	JobTitles  []string `json:"job titles,omitempty"`
}

func loadTestResumes(filename string) []Resume {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var resume Resume
	if err := json.Unmarshal(content, &resume); err != nil {
		log.Fatal(err)
	}
	return []Resume{resume}
}

func loadTestFilters(filename string) []Filter {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var filter Filter
	if err := json.Unmarshal(content, &filter); err != nil {
		log.Fatal(err)
	}
	return []Filter{filter}
}

func getResumes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, loadTestResumes("resume.json"))
}

func getFilters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, loadTestFilters("filter.json"))
}

func main() {
	fmt.Println("Setting up router")
	router := gin.Default()
	router.GET("/resumes", getResumes)
	router.GET("/filters", getFilters)
	router.Run("localhost:8080")
}
