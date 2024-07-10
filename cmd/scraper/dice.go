package cmd

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/lcsebastian/urban-palm-tree/cmd"
)

type Filter cmd.Filter
type Job cmd.Job

const (
	baseUrl = "https://dice.com"
)

// getQuery converts a Filter struct into a query string for dice.com.
// If the conversion is successful the string is returned, else nil is returned.
func getQuery(filter *Filter) string {
	return "I'm a stub"
}

func Scrape(filter Filter) []Job {
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
