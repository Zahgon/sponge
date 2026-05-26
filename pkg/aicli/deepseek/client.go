// Package deepseek provides a client for the DeepSeek API.
package deepseek

import (
	"github.com/go-dev-frame/sponge/pkg/aicli/chatgpt"
)

// https://api-docs.deepseek.com/

const (
	BaseURL = "https://api.deepseek.com/"

	ModelDeepSeekChat     = "deepseek-chat"
	ModelDeepSeekReasoner = "deepseek-reasoner"
	DefaultModel          = ModelDeepSeekChat

	TemperatureCodeGeneration float32 = 0.0 // code generation, mathematical problem-solving
	TemperatureDataAnalysis   float32 = 1.0 // data extraction, analysis
	TemperatureDataChat       float32 = 1.3 // universal chat, translation
	TemperatureDataCreative   float32 = 2.0 // creative writing, poetry writing
)

type (
	Client         = chatgpt.Client
	ClientOption   = chatgpt.ClientOption
	ContextMessage = chatgpt.ContextMessage
)

var (
	WithMaxTokens              = chatgpt.WithMaxTokens
	WithModel                  = chatgpt.WithModel
	WithTemperature            = chatgpt.WithTemperature
	WithInitialRole            = chatgpt.WithInitialRole
	WithInitialContextMessages = chatgpt.WithInitialContextMessages
	WithEnableContext          = chatgpt.WithEnableContext
)

// NewClient creates a new chat client.
func NewClient(apiKey string, opts ...ClientOption) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
