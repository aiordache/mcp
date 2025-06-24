package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (c *SOClient) GetAnswer(answerID int) (*Answer, error) {
	queryParam := url.Values{}
	queryParam.Add("order", "desc")
	queryParam.Add("sort", "activity")
	queryParam.Add("filter", "!*Mg4Pjg8kf1XZzrq") // Filter created to retrieve `body` in response.

	endpoint := fmt.Sprintf("/answers/%d", answerID)
	req, err := c.newRequest(http.MethodGet, endpoint, nil, queryParam)
	if err != nil {
		return nil, err
	}

	data, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	var resp AnswerResponse
	if err = json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to decode answer body response: %v", err)
	}
	if len(resp.Items) == 0 {
		return nil, fmt.Errorf("no answer found for ID %d", answerID)
	}
	return &resp.Items[0], nil
}
