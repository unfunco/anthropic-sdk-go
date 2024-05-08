// Package anthropic provides a client library for interacting with the
// Anthropic safety-first language model REST APIs.
package anthropic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultAPIVersion    = "2023-06-01"
	defaultBaseURL       = "https://api.anthropic.com/v1/"
	defaultBaseUserAgent = "unfunco/anthropic-sdk-go"
)

// Client manages communication with the Anthropic REST API.
type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
	reusable   service

	Messages *MessagesService
}

type service struct{ client *Client }

// NewClient returns a new Anthropic REST API client. If a nil httpClient is
// provided, a new http.Client will be used.
// Adapted from go-github's NewClient method:
// https://github.com/google/go-github/blob/master/github/github.go
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	claude := &Client{httpClient: httpClient}
	claude.baseURL, _ = url.Parse(defaultBaseURL)
	claude.reusable.client = claude

	claude.Messages = (*MessagesService)(&claude.reusable)

	return claude
}

// NewRequest creates an API request. A relative URL can be provided in path,
// in which case it is resolved relative to the BaseURL of the Client.
// Paths should always be specified without a preceding slash.
// If specified, the value pointed to by body is JSON encoded and included as
// the request body.
// Adapted from go-github's Client.NewRequest method:
// https://github.com/google/go-github/blob/master/github/github.go
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
	req.Header.Set("user-agent", defaultBaseUserAgent+"@"+semanticVersion)

	if body != nil {
		req.Header.Set("content-type", "application/json")
	}

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an error
// if an API error has occurred. If v implements the io.Writer interface, the
// raw response body will be written to v without attempting to first decode it.
// If v is nil, and no error occurs, the response is returned as-is.
// Adapted from go-github's Client.BareDo and Client.Do methods:
// https://github.com/google/go-github/blob/master/github/github.go
func (c *Client) Do(ctx context.Context, req *http.Request, v any) (*http.Response, error) {
	if ctx == nil {
		return nil, errors.New("context must be non-nil")
	}

	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
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

	return resp, err
}
