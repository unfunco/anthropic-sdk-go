package anthropic

import "fmt"

// Usage defines billing and rate-limit usage information. Billing and
// rate-limiting are driven by token counts.
type Usage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

// String implements the fmt.Stringer interface for Usage.
func (u *Usage) String() string {
	return fmt.Sprintf("Input: %d, Output: %d", u.InputTokens, u.OutputTokens)
}
