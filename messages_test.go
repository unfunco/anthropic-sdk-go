package anthropic

import (
	"testing"
)

func TestMessagesService_Create(t *testing.T) {
	_ = &CreateMessageOptions{
		MaxTokens: 1024,
	}
	t.SkipNow()
}
