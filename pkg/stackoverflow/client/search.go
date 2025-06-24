package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (c *SOClient) SearchQuestions(query string, page int) ([]SearchItem, error) {
	queryParam := url.Values{}
	queryParam.Add("order", "desc")
	queryParam.Add("sort", "activity")
	queryParam.Add("page", fmt.Sprintf("%d", page))
	queryParam.Add("intitle", query)
	req, err := c.newRequest(http.MethodGet, "/search", []byte{}, queryParam)
	if err != nil {
		return nil, err
	}

	data, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	var searchResponse SearchResponse
	if err = json.Unmarshal(data, &searchResponse); err != nil {
		return nil, fmt.Errorf("fail to decode search response %#v: %v", data, err)
	}

	return searchResponse.Items, nil
}

type SearchResponseWithAnswers struct {
	IsAnswered bool     `json:"is_answered"`
	QuestionID int      `json:"question_id"`
	Link       string   `json:"link"`
	Title      string   `json:"title"`
	Answers    []Answer `json:"answers_list"`
}

func (c *SOClient) SearchAnswers(query string, page int) ([]SearchResponseWithAnswers, error) {
	queryParam := url.Values{}
	queryParam.Add("order", "desc")
	queryParam.Add("sort", "activity")
	queryParam.Add("page", fmt.Sprintf("%d", page))
	queryParam.Add("intitle", query)
	req, err := c.newRequest(http.MethodGet, "/search", []byte{}, queryParam)
	if err != nil {
		return nil, err
	}

	data, err := c.sendRequest(req)
	if err != nil {
		return nil, err
	}

	var searchResponse SearchResponse
	if err = json.Unmarshal(data, &searchResponse); err != nil {
		return nil, fmt.Errorf("fail to decode search response %#v: %v", data, err)
	}

	// fetch all possible answers for every questionID we've received
	var output []SearchResponseWithAnswers
	for _, item := range searchResponse.Items {
		var result SearchResponseWithAnswers
		result.IsAnswered = item.IsAnswered
		result.QuestionID = item.QuestionID
		result.Link = item.Link
		result.Title = item.Title

		// fetch QuestionID
		questionID := item.QuestionID
		answers, err := c.GetAnswerIDsByQuestionID(questionID)
		if err != nil {
			continue
		}

		for _, ans := range answers {
			answer, err := c.GetAnswer(ans.AnswerID)
			if err != nil || answer == nil {
				continue
			}

			// append answers
			result.Answers = append(result.Answers, *answer)
		}

		output = append(output, result)
	}

	return output, nil
}
