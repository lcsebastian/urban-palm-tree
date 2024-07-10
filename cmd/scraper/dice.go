package cmd

import (
//	"github.com/gocolly/colly"
"github.com/lcsebastian/urban-palm-tree/cmd"
)

type Filter cmd.Filter

type SalaryRange struct {
	Min float64
	Max float64
}

type JobType int

const (
	FullTime JobType = iota
	PartTime
	Contract
)

type Job struct {
	Link        string
	Title       string
	Description string
	Company     string
	Location    string
	Remote      bool
	Salary      SalaryRange
	Type        JobType
	PostedDate  string
	Skills      []string
}

const (
	baseUrl = "https://dice.com"
)

func Scrape(filter Filter) []Job {
	return nil
}
