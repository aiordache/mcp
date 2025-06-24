package client_test

import (
	"testing"

	"github.com/algolia/mcp/pkg/stackoverflow/client"
	"github.com/stretchr/testify/require"
)

func TestTags(t *testing.T) {
	c := client.NewClient(`<token>`)
	tags, err := c.ListTags()
	require.NoError(t, err)
	require.NotEmpty(t, tags)
}

func TestSearch(t *testing.T) {
	c := client.NewClient(`<token>`)
	response, err := c.Search("composition", 1)
	require.NoError(t, err)
	require.NotEmpty(t, response)
}
