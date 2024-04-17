package anthropic

import (
	"context"
	"net/http"
)

type MessagesService service

// Message defines ...
type Message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

// Usage defines billing and rate-limit usage information.
// Billing and rate-limiting are driven by token counts, since tokens represent
// the underlying cost to Anthropic.
type Usage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

type Content struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// CreateMessageInput defines a structured list of input messages.
type CreateMessageInput struct {
	// MaxTokens defines the maximum number of tokens to generate before
	// stopping. Token generation may stop before reaching this limit, this only
	// specifies the absolute maximum number of tokens to generate. Different
	// models have different maximum token limits.
	MaxTokens int `json:"max_tokens"`
	// Messages are the input messages, models are trained to operate on
	// alternating user and assistant conversational turns. When creating a new
	// message, prior conversational turns can be specified with this field.
	Messages []Message `json:"messages"`
	// Model defines the language model that will be used to complete the
	// prompt. See model.go for a list of available models.
	Model LanguageModel `json:"model"`
	// StopSequences defines custom text sequences that will cause the model to
	// stop generating.
	StopSequences []string `json:"stop_sequences,omitempty"`
	// System provides a means of specifying context and instructions to the
	// model, such as specifying a particular goal or role.
	System string `json:"system,omitempty"`
	// Temperature defines the amount of randomness injected into the response.
	// Note that even with a temperature of 0.0, results will not be fully
	// deterministic.
	Temperature *float64 `json:"temperature,omitempty"`
	// TopK is used to remove long tail low probability responses.
	// Recommended for advanced use cases only. You usually only need to use
	// Temperature.
	TopK *int `json:"top_k,omitempty"`
	// TopP (nucleus-sampling) defines the cumulative probability of the highest probability.
	// Recommended for advanced use cases only. You usually only need to use
	// Temperature.
	TopP *float64 `json:"top_p,omitempty"`
}

type CreateMessageOutput struct {
	Id           *string    `json:"id"`
	Type         *string    `json:"type"`
	Role         *string    `json:"role"`
	Model        *string    `json:"model"`
	StopSequence *string    `json:"stop_sequence"`
	StopReason   *string    `json:"stop_reason"`
	Content      []*Content `json:"content"`
	Usage        *Usage     `json:"usage"`
}

func (c *CreateMessageOutput) String() string {
	return c.Content[0].Text
}

// Create creates a new message using the provided options.
func (c *MessagesService) Create(
	ctx context.Context,
	in *CreateMessageInput,
) (*CreateMessageOutput, *http.Response, error) {
	req, err := c.client.NewRequest(http.MethodPost, "messages", in)
	if err != nil {
		return nil, nil, err
	}

	out := new(CreateMessageOutput)
	resp, err := c.client.Do(ctx, req, out)
	if err != nil {
		return nil, resp, err
	}

	return out, resp, nil
}
