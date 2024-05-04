package anthropic

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

// StreamMessageOptions defines the options available when streaming server-sent
// events from the Anthropic REST API, the StreamMessageOptions definition is
// currently identical to the CreateMessageOptions definition, but the
// StreamMessageOptions type has a custom MarshalJSON implementation that will
// append a stream field and set it to true.
type StreamMessageOptions struct {
	// Temperature defines the amount of randomness injected into the response.
	// Note that even with a temperature of 0.0, results will not be fully
	// deterministic.
	Temperature *float64 `json:"temperature,omitempty"`
	// TopK is used to remove long tail low probability responses by only
	// sampling from the top K options for each subsequent token.
	// Recommended for advanced use cases only. You usually only need to use
	// Temperature.
	TopK *int `json:"top_k,omitempty"`
	// TopP is the nucleus-sampling parameter. Temperature or TopP should be
	// used, but not both.
	// Recommended for advanced use cases only. You usually only need to use
	// Temperature.
	TopP *float64 `json:"top_p,omitempty"`
	// Model defines the language model that will be used to complete the
	// prompt. See model.go for a list of available models.
	Model LanguageModel `json:"model"`
	// System provides a means of specifying context and instructions to the
	// model, such as specifying a particular goal or role.
	System string `json:"system,omitempty"`
	// Messages are the input messages, models are trained to operate on
	// alternating user and assistant conversational turns. When creating a new
	// message, prior conversational turns can be specified with this field,
	// and the model generates the next Message in the conversation.
	Messages []Message `json:"messages"`
	// StopSequences defines custom text sequences that will cause the model to
	// stop generating. If the model encounters any of the sequences, the
	// StopReason field will be set to "stop_sequence" and the response
	// StopSequence field will be set to the sequence that caused the model to
	// stop.
	StopSequences []string `json:"stop_sequences,omitempty"`
	// MaxTokens defines the maximum number of tokens to generate before
	// stopping. Token generation may stop before reaching this limit, this only
	// specifies the absolute maximum number of tokens to generate. Different
	// models have different maximum token limits.
	MaxTokens int `json:"max_tokens"`
}

// MarshalJSON implements the json.Marshaler interface for StreamMessageOptions.
// When StreamMessageOptions is marshalled to JSON, a stream field will be added
// and set to a boolean value of true.
func (c *StreamMessageOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Temperature   *float64      `json:"temperature,omitempty"`
		TopK          *int          `json:"top_k,omitempty"`
		TopP          *float64      `json:"top_p,omitempty"`
		Model         LanguageModel `json:"model,omitempty"`
		System        string        `json:"system,omitempty"`
		Messages      []Message     `json:"messages,omitempty"`
		StopSequences []string      `json:"stop_sequences,omitempty"`
		MaxTokens     int           `json:"max_tokens,omitempty"`
		Stream        bool          `json:"stream"`
	}{
		Temperature:   c.Temperature,
		TopK:          c.TopK,
		TopP:          c.TopP,
		Model:         c.Model,
		System:        c.System,
		Messages:      c.Messages,
		StopSequences: c.StopSequences,
		MaxTokens:     c.MaxTokens,
		Stream:        true,
	})
}

// ServerSentEvent defines a server-sent event.
type ServerSentEvent struct {
	Event *string
	Data  string
	Raw   []string
}

// Stream creates a new message using the provided options and streams the
// response using server-sent events. This is a convenience method that
// combines the Create and Stream methods.
func (c *MessagesService) Stream(
	ctx context.Context,
	opts *StreamMessageOptions,
) (*<-chan ServerSentEvent, *http.Response, error) {
	req, err := c.client.NewRequest(http.MethodPost, "messages", opts)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, nil)
	if err != nil {
		return nil, resp, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	output, err := newServerSentEventStream(resp.Body)

	return output, resp, err
}

func newServerSentEventStream(body io.ReadCloser) (*<-chan ServerSentEvent, error) {
	scanner := bufio.NewScanner(body)
	scanner.Buffer(make([]byte, 4096), bufio.MaxScanTokenSize)
	scanner.Split(func(data []byte, atEOF bool) (int, []byte, error) {
		return 0, nil, nil
	})

	// TODO

	return new(<-chan ServerSentEvent), nil
}
