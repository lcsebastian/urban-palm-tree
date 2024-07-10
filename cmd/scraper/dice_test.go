package cmd

import (
	"github.com/lcsebastian/urban-palm-tree/cmd"
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
)	

func TestScrape(t *testing.T) {
	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")
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
