// Package anthropic provides a client library for interacting with the
// Anthropic safety-first language model REST APIs.
package anthropic

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultAPIVersion = "2023-06-01"
	defaultBaseURL    = "https://api.anthropic.com/v1/"
	defaultUserAgent  = "unfunco/anthropic-sdk-go@v0.0.0"
)

// Client manages communication with the Anthropic REST API.
type Client struct {
	baseURL  *url.URL
	client   *http.Client
	reusable service

	Messages *MessagesService
}

type service struct{ client *Client }

// NewClient returns a new Anthropic REST API client.
func NewClient(client *http.Client) *Client {
	if client == nil {
		client = &http.Client{}
	}

	claude := &Client{client: client}
	claude.baseURL, _ = url.Parse(defaultBaseURL)
	claude.reusable.client = claude
	claude.Messages = (*MessagesService)(&claude.reusable)

	return claude
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, path string, body any) (*http.Request, error) {
	u, err := c.baseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err = enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("anthropic-version", defaultAPIVersion)
	req.Header.Set("user-agent", defaultUserAgent)

	if body != nil {
		req.Header.Set("content-type", "application/json")
	}

	return req, nil
}

// Do sends an API request and returns the API response.
func (c *Client) Do(ctx context.Context, req *http.Request, v any) (*http.Response, error) {
	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil
		}
		if decErr != nil {
			err = decErr
		}
	}

	return resp, nil
}
