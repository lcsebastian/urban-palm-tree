package cmd

import (
	"fmt"
	"testing"

	"github.com/lcsebastian/urban-palm-tree/cmd"
	"github.com/stretchr/testify/assert"
)

func TestScrape(t *testing.T) {
	filter := cmd.LoadTestFilters("../../data/filter.json")[0]
	assert.NotNil(t, filter)
	Scrape(filter)
}

func TestGetQuery(t *testing.T) {
	filter := cmd.LoadTestFilters("../../data/filter.json")[0]
	assert.NotNil(t, filter)

	expectedQuery := DiceQuery{
		[]string{"software engineer", "java", "full stack", "gcp", "jenkins"},
		"Los Angeles, CA, USA",
		[]string{"Remote"},
		[]string{"FULLTIME", "CONTRACTS"},
		"ONE",
	}
	actualQuery := getQuery(filter)
	assert.NotNil(t, actualQuery)
	assert.Equal(t, expectedQuery, actualQuery)
	fmt.Println("Stringer'd query : ", actualQuery)
}
