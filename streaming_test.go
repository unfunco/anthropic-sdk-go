package anthropic

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestStreamMessageOptions_MarshalJSON_Empty(t *testing.T) {
	options, err := json.Marshal(&StreamMessageOptions{})
	if err != nil {
		t.Error(err)
	}

	expected := []byte(`{"stream":true}`)
	if !bytes.Equal(options, expected) {
		t.Errorf("options = %q, want %q", options, expected)
	}
}

func TestStreamMessageOptions_MarshalJSON_Initialised(t *testing.T) {
	options, err := json.Marshal(&StreamMessageOptions{
		MaxTokens: 512,
		Messages: []Message{
			{
				Content: "This is a test message.",
				Role:    "user",
			},
		},
		Model:         Claude3Opus20240229,
		StopSequences: make([]string, 0),
		System:        "",
	})
	if err != nil {
		t.Error(err)
	}

	expected := []byte(`{"model":"claude-3-opus-20240229","messages":[{"content":"This is a test message.","role":"user"}],"max_tokens":512,"stream":true}`)
	if !bytes.Equal(options, expected) {
		t.Errorf("options = %q, want %q", options, expected)
	}
}

func TestMessagesService_Stream(t *testing.T) {
	bytes.NewBufferString(`
event: message_start
data: {"type": "message_start", "message": {"id": "msg_1nZdL29xx5MUA1yADyHTEsnR8uuvGzszyY", "type": "message", "role": "assistant", "content": [], "model": "claude-3-opus-20240229", "stop_reason": null, "stop_sequence": null, "usage": {"input_tokens": 25, "output_tokens": 1}}}

event: content_block_start
data: {"type": "content_block_start", "index": 0, "content_block": {"type": "text", "text": ""}}

event: ping
data: {"type": "ping"}

event: content_block_delta
data: {"type": "content_block_delta", "index": 0, "delta": {"type": "text_delta", "text": "Hello"}}

event: content_block_delta
data: {"type": "content_block_delta", "index": 0, "delta": {"type": "text_delta", "text": "!"}}

event: content_block_stop
data: {"type": "content_block_stop", "index": 0}

event: message_delta
data: {"type": "message_delta", "delta": {"stop_reason": "end_turn", "stop_sequence":null}, "usage": {"output_tokens": 15}}

event: message_stop
data: {"type": "message_stop"}
`)
	t.SkipNow()
}
