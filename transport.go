package anthropic

import (
	"errors"
	"net/http"
	"net/url"
)

// Transport ...
type Transport struct {
	APIKey string
}

func (t *Transport) Client() *http.Client {
	return &http.Client{
		Transport: t,
	}
}

// RoundTrip implements the http.RoundTripper interface.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.APIKey == "" {
		return nil, errors.New("API key is required")
	}

	newReq := new(http.Request)
	*newReq = *req
	newReq.URL = new(url.URL)
	*newReq.URL = *req.URL

	newReq.Header = make(http.Header, len(req.Header))
	for k, v := range req.Header {
		newReq.Header[k] = v
	}

	newReq.Header.Set("x-api-key", t.APIKey)

	return http.DefaultTransport.RoundTrip(newReq)
}
