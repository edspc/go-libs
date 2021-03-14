package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
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

var (
	client BaseClient
)

func init() {
	client.HTTPClient = &http.Client{Timeout: time.Minute}
	client.Header = &http.Header{}
}

func SetBaseURL(url string) {
	client.BaseURL = url
}

func (c BaseClient) GetUrl(url string) string {
	return c.BaseURL + "/" + strings.TrimLeft(url, "/")
}

func SetAuthHeader(authType string, authTocken string) {
	client.Header.Set("Authorization", authType+" "+authTocken)
}

func AddHeader(name string, value string) {
	client.Header.Add(name, value)
}

func (c *BaseClient) SendRequest(req *http.Request, v interface{}) (*http.Response, error) {
	req.Header = *client.Header
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
