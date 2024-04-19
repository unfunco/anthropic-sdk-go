package anthropic

// LanguageModel represents a language model that can be used to complete
// a prompt.
// https://docs.anthropic.com/claude/docs/models-overview
type LanguageModel string

const (
	Claude3Opus20240229   LanguageModel = "claude-3-opus-20240229"
	Claude3Sonnet20240229 LanguageModel = "claude-3-sonnet-20240229"
	Claude3Haiku20240307  LanguageModel = "claude-3-haiku-20240307"
	Claude21              LanguageModel = "claude-2.1"
	Claude20              LanguageModel = "claude-2.0"
	ClaudeInstant12       LanguageModel = "claude-instant-1.2"
)
