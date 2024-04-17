package anthropic

import (
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	c1 := NewClient(nil)
	c2 := NewClient(nil)

	if c1.httpClient == c2.httpClient {
		t.Error("NewClient returned equal http.Clients but they should differ")
	}
}

func TestClient_NewRequest(t *testing.T) {
	c := NewClient(nil)
	req, err := c.NewRequest(http.MethodPost, "example", nil)
	if err != nil {
		t.Error(err)
	}

	if req.Method != http.MethodPost {
		t.Errorf("req.Method = %q, want %q", req.Method, http.MethodPost)
	}

	if req.URL.String() != "https://api.anthropic.com/v1/example" {
		t.Errorf("req.URL = %q, want %q", req.URL.String(), "https://api.anthropic.com/v1/example")
	}
}
