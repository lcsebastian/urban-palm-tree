package cmd

import (
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"

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

func getAbsolutePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return dir
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

func formatJobTitle(title string) string {
	title = strings.TrimSpace(title)
	title = strings.ReplaceAll(title, "  ", "")
	title = strings.ReplaceAll(title, "\n", " ")
	return title
}

// Cannot use colly since Dice is so full of javascript and AJAX
func CollyScrape() []Job {
	jobs := []Job{}
	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
	// Instantiate default collector
	c := colly.NewCollector()
	c.WithTransport(t)

	// On every a element which has href attribute call callback
	c.OnHTML(".card-title-link", func(e *colly.HTMLElement) {
		if e.Attr("class") == "card-title-link normal" {
			id := e.Attr("id")
			link := fmt.Sprintf("%v/job-detail/%v", getBaseUrl(), id)
			jobs = append(jobs, Job{
				Link:  link,
				Title: formatJobTitle(e.Text),
			})
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	c.Visit("file://" + getAbsolutePath() + "/data/dice-page.html")
	c.Wait()

	return jobs
}
