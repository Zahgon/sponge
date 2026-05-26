// Package chatgpt provides a client for the OpenAI chat GPT API.
package chatgpt

import (
	"context"

	"github.com/sashabaranov/go-openai"

	"github.com/go-dev-frame/sponge/pkg/aicli"
)

// https://platform.openai.com/docs/api-reference

// Client is a chat GPT client.
type Client struct {
	apiKey    string
	maxTokens int
	ModelName string
	Cli       *openai.Client

	/*
		| Temperature Value | Randomness   | Applicable Scenarios                      |
		|----------------|------------------------|-------------------------------------------|
		| 0                 | No randomness         | Factual answers, code generation, technical documentation |
		| 0.5 - 0.7      | Moderate randomness | Conversational systems, content creation, recommendation systems |
		| 1                 | High randomness       | Creative writing, brainstorming, advertising copy |
		| 1.5 - 2         | Extreme randomness  | Artistic creation, game design, exploratory tasks |
	*/
	temperature float32

	roleDesc string // initial role description

	enableContext   bool                           // whether to use assistant context, default is false
	contextMessages []openai.ChatCompletionMessage // initial context messages
}

// NewClient creates a new chat client.
func NewClient(apiKey string, opts ...ClientOption) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Send sends a prompt to the chat gpt and returns the response.
func (c *Client) Send(ctx context.Context, prompt string, files ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Deprecated

// SendStream sends a prompt to the chat gpt and returns a channel of responses.
func (c *Client) SendStream(ctx context.Context, prompt string, files ...string) *aicli.StreamReply {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated

//nolint

// ListModelNames lists all available model names.
func (c *Client) ListModelNames(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListContextMessages list assistant context messages
func (c *Client) ListContextMessages() []*ContextMessage { _ = "STUB: not implemented"; return nil }

// RefreshContext refreshes assistant context
func (c *Client) RefreshContext() { _ = "STUB: not implemented"; return }

// ModifyInitialRole modifies the initial role description.
func (c *Client) ModifyInitialRole(roleDesc string) { _ = "STUB: not implemented"; return }

func (c *Client) appendAssistantContext(prompt string, replyContent string) {
	_ = "STUB: not implemented"
	return
}

func (c *Client) setMessages(ctx context.Context, prompt string, files ...string) ([]openai.ChatCompletionMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// history context

// file message

// user message

func (c *Client) uploadFiles(ctx context.Context, files []string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// for assistants
