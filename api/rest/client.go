package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type BaseClient struct {
	BaseURL    string
	Header     *http.Header
	HTTPClient HTTPClient
}

type errorResponse struct {
	Message string `json:"message"`
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func (c *BaseClient) SetBaseURL(url string) {
	c.BaseURL = url
}

func (c *BaseClient) GetUrl(url string) string {
	return c.BaseURL + "/" + strings.TrimLeft(url, "/")
}

func (c *BaseClient) GetURLWithQuery(query url.Values) string {
	return c.BaseURL + "?" + query.Encode()
}

func (c *BaseClient) SetAuthHeader(authType string, authToken string) {
	c.Header.Set("Authorization", authType+" "+authToken)
}

func (c *BaseClient) AddHeader(name string, value string) {
	c.Header.Add(name, value)
}

func (c *BaseClient) SendRequest(req *http.Request, v interface{}) (*http.Response, error) {
	req.Header = *c.Header
	req.Header.Set("Accept", "application/json; charset=utf-8")
	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return nil, errors.New(errRes.Message)
		}

		return nil, fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return nil, err
	}

	return res, nil
}
