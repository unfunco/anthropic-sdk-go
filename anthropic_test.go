package anthropic

import "testing"

func TestNewClient(t *testing.T) {
	c1 := NewClient(nil)
	c2 := NewClient(nil)

	if c1.client == c2.client {
		t.Error("NewClient returned equal http.Clients but they should differ")
	}
}

func TestClient_NewRequest(t *testing.T) {
	_ = NewClient(nil)
}
