package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-retryablehttp"
)

type SOClient struct {
	client        *http.Client
	baseURL       string
	apiAcessToken string
	team          string
}

const (
	BaseURL = "https://api.stackoverflowteams.com/2.3"
	Team    = "algolia"
)

func NewClient(apiAcessToken string) *SOClient {
	withRetry := retryablehttp.NewClient()
	return &SOClient{
		client:        withRetry.StandardClient(),
		baseURL:       BaseURL,
		apiAcessToken: apiAcessToken,
		team:          Team,
	}
}

func (c *SOClient) newRequest(method, path string, payload []byte, urlParams url.Values) (*http.Request, error) {
	req, err := http.NewRequest(method, c.baseURL+path, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	urlParams.Add("team", c.team)
	req.URL.RawQuery = urlParams.Encode()
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-API-Access-Token", c.apiAcessToken)

	return req, nil
}

func (c *SOClient) sendRequest(req *http.Request) ([]byte, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // nolint: errcheck

	if resp.StatusCode != http.StatusOK {
		var soErr SOError
		if err = json.NewDecoder(resp.Body).Decode(&soErr); err != nil {
			return nil, fmt.Errorf("fail to decode so error - %s: %d, %s", err, resp.StatusCode, resp.Body)
		}
		return nil, &soErr
	}
	answer, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return answer, nil
}
