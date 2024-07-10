package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

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

func scrape_jobs(filter Filter) []Job {
	return nil
}
