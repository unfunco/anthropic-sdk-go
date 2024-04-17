package anthropic

import (
	"errors"
	"net/http"
	"net/url"
)

// Transport is a http.RoundTripper that includes an API key in each request.
type Transport struct {
	APIKey string
}

// Client returns an HTTP client that will include the API key in the request,
// and is safe for concurrent use by multiple goroutines.
func (t *Transport) Client() *http.Client {
	return &http.Client{
		Transport: t,
	}
}

// RoundTrip implements the http.RoundTripper interface.
// It makes a copy of the HTTP request so that it complies with the requirements
// of the interface and adds the API key to the new request before calling the
// default http.RoundTripper.
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
