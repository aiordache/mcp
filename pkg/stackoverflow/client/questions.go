package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type AnswerOwner struct {
	DisplayName string `json:"display_name"`
	Link        string `json:"link"`
}

type AnswerItem struct {
	AnswerID int         `json:"answer_id"`
	Owner    AnswerOwner `json:"owner"`
}

type AnswersResponse struct {
	Items []AnswerItem `json:"items"`
}

func (c *SOClient) GetAnswerIDsByQuestionID(questionID int) ([]AnswerItem, error) {
	queryParam := url.Values{}
	queryParam.Add("order", "desc")
	queryParam.Add("sort", "activity")

	endpoint := fmt.Sprintf("/questions/%d/answers", questionID)

	req, err := c.newRequest(http.MethodGet, endpoint, []byte{}, queryParam)
	if err != nil {
		return nil, fmt.Errorf("fail newRequest: %w", err)
	}

	data, err := c.sendRequest(req)
	if err != nil {
		return nil, fmt.Errorf("fail sendRequest: %w", err)
	}

	var resp AnswersResponse
	if err = json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("fail to decode answers response: %v", err)
	}

	return resp.Items, nil
}
