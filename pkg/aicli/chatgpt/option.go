package chatgpt

import "github.com/sashabaranov/go-openai"

const (
	ModelGPT3Dot5Turbo = openai.GPT3Dot5Turbo
	ModelGPT4          = openai.GPT4
	ModelGPT4Turbo     = openai.GPT4Turbo
	ModelGPT4o         = openai.GPT4o // default
	ModelGPT4oMini     = openai.GPT4oMini
	ModelO1Mini        = openai.O1Mini
	ModelO1Preview     = openai.O1Preview

	DefaultModel     = ModelGPT4o
	defaultMaxTokens = 8192
)

// ClientOption is a function that sets a Client option.
type ClientOption func(*Client)

func defaultClientOptions() *Client { _ = "STUB: not implemented"; return nil }

// default is false

func (c *Client) apply(opts ...ClientOption) { _ = "STUB: not implemented"; return }

// WithMaxTokens sets the maximum number of tokens
func WithMaxTokens(maxTokens int) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

// WithModel sets the model name
func WithModel(name string) ClientOption { _ = "STUB: not implemented"; return *new(ClientOption) }

// WithTemperature sets the temperature
func WithTemperature(temperature float32) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

// WithInitialRole sets the initial role type
func WithInitialRole(roleDesc string) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

// WithEnableContext sets assistant context
func WithEnableContext() ClientOption { _ = "STUB: not implemented"; return *new(ClientOption) }

// ContextMessage chat history message
type ContextMessage struct {
	Role    string `json:"role"` // system, user, assistant, etc.
	Content string `json:"content"`
}

// WithInitialContextMessages sets initial context messages, automatically set enableContext to true
func WithInitialContextMessages(messages ...*ContextMessage) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}
