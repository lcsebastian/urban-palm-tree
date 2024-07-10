package cmd

import (
	"encoding/json"
	"log"
	"os"
	"fmt"
)

type JobType string

const (
	FullTime, PartTime, Contract JobType = "FullTime", "PartTime", "Contract"
)

type PostedDateEnum string

const (
	Any, Today PostedDateEnum = "Any", "Today"
)

type SalaryRange struct {
	Min float64
	Max float64
}

type LocationInfo struct {
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	Country    string `json:"country,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
}

func (l LocationInfo) String() string {
	return fmt.Sprintf("%v, %v, %v", l.City, l.State, l.Country)
}

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
		Location LocationInfo `json:"location,omitempty"`
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
	ID         int            `json:"id,omitempty"`
	PostedDate PostedDateEnum `json:"date posted,omitempty"`
	RemoteOnly bool           `json:"remote only,omitempty"`
	Location   LocationInfo `json:"location,omitempty"`
	Type       []JobType      `json:"job type,omitempty"`
	MinimumPay float64        `json:"minimum pay,omitempty"`
	Keywords   []string       `json:"keywords,omitempty"`
	JobTitles  []string       `json:"job titles,omitempty"`
}

type Job struct {
	Link        string
	Title       string
	Description string
	Company     string
	Location    LocationInfo
	Remote      bool
	Salary      SalaryRange
	Type        JobType
	PostedDate  PostedDateEnum
	Skills      []string
}

func LoadTestResumes(filename string) []Resume {
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

func LoadTestFilters(filename string) []Filter {
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
