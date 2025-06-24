package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type listTagsResponse struct {
	Response
	Items []TagItem `json:"items"`
}

func (c *SOClient) ListTags() ([]TagItem, error) {
	var tagList []TagItem
	page := 1

	for {
		queryParam := url.Values{}
		queryParam.Add("page", strconv.Itoa(page))
		req, err := c.newRequest(http.MethodGet, "/tags", []byte{}, queryParam)
		if err != nil {
			return nil, err
		}
		data, err := c.sendRequest(req)
		if err != nil {
			return nil, err
		}

		var listTagsResponse listTagsResponse
		if err = json.Unmarshal(data, &listTagsResponse); err != nil {
			return nil, fmt.Errorf("fail to decodes tags %#v: %v", data, err)
		}
		tagList = append(tagList, listTagsResponse.Items...)
		if !listTagsResponse.HasMore {
			break
		}
		page++
	}
	return tagList, nil
}
