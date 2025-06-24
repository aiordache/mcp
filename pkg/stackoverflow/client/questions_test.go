package client

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAnswersIDs(t *testing.T) {
	token := os.Getenv("STACKOVERFLOW_API_KEY")
	c := NewClient(token)
	response, err := c.GetAnswerIDsByQuestionID(9019)
	t.Logf("response is %v", response)
	require.NoError(t, err)
	require.NotEmpty(t, response)
}
