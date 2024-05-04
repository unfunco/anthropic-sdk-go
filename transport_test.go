package anthropic

import "testing"

func TestNewTransport(t *testing.T) {
	apiKey := "test-api-key"
	transport := NewTransport(apiKey)
	if transport.APIKey != apiKey {
		t.Errorf("transport.APIKey = %q, want %q", transport.APIKey, apiKey)
	}
}
