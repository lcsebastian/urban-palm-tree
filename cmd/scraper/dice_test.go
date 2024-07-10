package cmd

import (
  "testing"
  "github.com/stretchr/testify/assert"
)



func TestScrape(t *testing.T) {
	// assert equality
  assert.Equal(t, 123, 123, "they should be equal")

  // assert inequality
  assert.NotEqual(t, 123, 456, "they should not be equal")
}

func TestGetQuery(t *testing.T) {
	filter := Filter {
		1,
		"TODAY",
		true,
		["Los Angeles, CA, USA"],
		
	}
}
