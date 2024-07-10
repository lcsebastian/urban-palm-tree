package cmd

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/lcsebastian/urban-palm-tree/cmd"
)

type Job cmd.Job

const (
	baseUrl = "https://dice.com"
)

// Defaults and Lookup Tables
func getDefaultWorkplaceTypes() []string {
	return []string{"Remote", "Hybrid"}
}

func jobTypeToEmploymentType(jobType cmd.JobType) string {
	switch jobType {
	case cmd.FullTime:
		return "FULLTIME"
	case cmd.PartTime:
		return "PARTTIME"
	case cmd.Contract:
		return "CONTRACTS"
	default:
		return ""
	}
}

func getPostedDate(postedDate cmd.PostedDateEnum) string {
	switch postedDate {
	case cmd.Today:
		return "ONE"
	default:
		return ""
	}
}

type DiceQuery struct {
	BaseQuery      []string
	Location       string
	WorkplaceType  []string
	EmploymentType []string
	PostedDate     string
}

func (d DiceQuery) String() string {
	return fmt.Sprintf("jobs?q=%v&location=%v&filters.postedDate=%v&filters.workplaceTypes=%v&filters.employmentType=%v&language=en",
		strings.Join(d.BaseQuery, " "),
		d.Location,
		d.PostedDate,
		strings.Join(d.WorkplaceType, "%7C"),
		strings.Join(d.EmploymentType, "%7C"))
}

// getQuery converts a Filter struct into a query object for dice.com.
func getQuery(filter cmd.Filter) DiceQuery {
	// build BaseQuery from job title + keywords
	result := DiceQuery{}
	result.BaseQuery = append(result.BaseQuery, append(filter.JobTitles, filter.Keywords...)...)

	result.Location = fmt.Sprint(filter.Location)

	result.WorkplaceType = getDefaultWorkplaceTypes()
	if filter.RemoteOnly {
		result.WorkplaceType = []string{"Remote"}
	}

	for _, value := range filter.Type {
		result.EmploymentType = append(result.EmploymentType, jobTypeToEmploymentType(value))
	}

	result.PostedDate = getPostedDate(filter.PostedDate)

	return result
}

func Scrape(filter cmd.Filter) []Job {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("dice.com", "www.dice.com"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(baseUrl)

	return nil
}
