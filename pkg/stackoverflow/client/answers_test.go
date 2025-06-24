package client

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAnswer(t *testing.T) {
	token := os.Getenv("STACKOVERFLOW_API_KEY")
	c := NewClient(token)
	response, err := c.GetAnswer(9023)
	t.Logf("response is %v", response)
	require.NoError(t, err)
	require.NotEmpty(t, response)
}
