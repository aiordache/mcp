package client_test

import (
	"os"
	"testing"

	"github.com/algolia/mcp/pkg/stackoverflow/client"
	"github.com/stretchr/testify/require"
)

func TestTags(t *testing.T) {
	token := os.Getenv("STACKOVERFLOW_API_KEY")
	c := client.NewClient(token)
	tags, err := c.ListTags()
	require.NoError(t, err)
	require.NotEmpty(t, tags)
}

func TestSearch(t *testing.T) {
	token := os.Getenv("STACKOVERFLOW_API_KEY")
	c := client.NewClient(token)
	response, err := c.SearchQuestions("composition", 1)
	require.NoError(t, err)
	require.NotEmpty(t, response)
}

func TestSearchAnswers(t *testing.T) {
	token := os.Getenv("STACKOVERFLOW_API_KEY")
	c := client.NewClient(token)
	response, err := c.SearchAnswers("composition", 1)
	t.Logf("%v", response)
	require.NoError(t, err)
	require.NotEmpty(t, response)
}
