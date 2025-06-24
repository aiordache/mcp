package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (c *SOClient) Search(query string, page int) ([]SearchItem, error) {
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
