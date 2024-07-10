package cmd

import (
	"fmt"
	"net/url"
	"strings"

	"context"
	"time"

	"os"

	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly"
	"github.com/lcsebastian/urban-palm-tree/cmd"
)

type Job cmd.Job

func getBaseUrl() string {
	return "https://dice.com"
}

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
	queryValues := url.Values{}
	queryValues.Add("q", strings.Join(d.BaseQuery, " "))
	queryValues.Add("location", d.Location)
	queryValues.Add("postedDate", d.PostedDate)
	queryValues.Add("workplaceTypes", strings.Join(d.WorkplaceType, "|"))
	queryValues.Add("employmentType", strings.Join(d.EmploymentType, "|"))
	queryValues.Add("language", "en")
	return fmt.Sprintf("jobs?%v", queryValues.Encode())
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

// chromedp works >:3
func Scrape(filter cmd.Filter) {
	fullUrl := fmt.Sprintf("%v/%v", getBaseUrl(), getQuery(filter))
	// Create a new context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Create a timeout context
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Variables to capture the scraped HTML
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(fullUrl),
		chromedp.Sleep(2*time.Second), // Wait for the page to load
		chromedp.OuterHTML("html", &res),
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Save HTML to file
	if err := os.WriteFile("../../data/dice-page.html", []byte(res), 0777); err != nil {
		fmt.Println("Error:", err)
		return
	}
}

// Cannot use colly since Dice is so full of javascript and AJAX
func CollyScrape(filter cmd.Filter) []Job {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("dice.com", "www.dice.com"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64; rv:128.0) Gecko/20100101 Firefox/128.0"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[class]", func(e *colly.HTMLElement) {
		class := e.Attr("class")
		if class == "card-title-link normal" {
			fmt.Printf("Class found: %q -> %s\n", e.Text, class)
			id := e.Attr("id")
			fmt.Printf("Id found: %s", id)
			link := fmt.Sprintf("%v/job-detail/%v", getBaseUrl(), id)
			// Visit link found on page
			// Only those links are visited which are in AllowedDomains
			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		// Print the response body (HTML)
		fmt.Println(string(r.Body))
	})

	fullUrl := fmt.Sprintf("%v/%v", getBaseUrl(), getQuery(filter))
	c.Visit(fullUrl)

	return nil
}
