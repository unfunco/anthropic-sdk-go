package anthropic

import (
	"context"
	"net/http"
)

type MessagesService service

type Message struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

// CreateMessageOptions ...
type CreateMessageOptions struct {
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

// CreateMessageOutput defines the response from creating a new message.
type CreateMessageOutput struct {
	ID           *string `json:"id"`
	Type         *string `json:"type"`
	Role         *string `json:"role"`
	Model        *string `json:"model"`
	StopSequence *string `json:"stop_sequence"`
	StopReason   *string `json:"stop_reason"`
	Usage        *Usage  `json:"usage"`
	Content      []*struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"content"`
}

// String implements the fmt.Stringer interface for CreateMessageOutput.
func (c *CreateMessageOutput) String() string {
	return c.Content[0].Text
}

// Create creates a new message using the provided options.
func (c *MessagesService) Create(
	ctx context.Context,
	opts *CreateMessageOptions,
) (*CreateMessageOutput, *http.Response, error) {
	req, err := c.client.NewRequest(http.MethodPost, "messages", opts)
	if err != nil {
		return nil, nil, err
	}

	output := new(CreateMessageOutput)
	resp, err := c.client.Do(ctx, req, output)
	if err != nil {
		return nil, resp, err
	}

	return output, resp, nil
}
