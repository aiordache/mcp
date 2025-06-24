package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

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

	var resp AnswerIDsResponse
	if err = json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("fail to decode answers response: %v", err)
	}

	return resp.Items, nil
}

func (c *SOClient) GetAnswers(questionID int) ([]Answer, error) {
	aIDs, err := c.GetAnswerIDsByQuestionID(questionID)
	if err != nil {
		return nil, err
	}
	if len(aIDs) == 0 {
		return nil, fmt.Errorf("no questions found")
	}

	// fetch all possible answers for every questionID we've received
	var output []Answer
	for _, item := range aIDs {
		answer, err := c.GetAnswer(item.AnswerID)
		if err != nil || answer == nil {
			continue
		}

		// append answers
		output = append(output, *answer)
	}

	return output, nil
}
