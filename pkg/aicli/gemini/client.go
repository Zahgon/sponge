// Package gemini provides a client for the Google generative AI API.
package gemini

import (
	"context"

	"github.com/google/generative-ai-go/genai"

	"github.com/go-dev-frame/sponge/pkg/aicli"
)

// Client is a Google generative AI client.
type Client struct {
	apiKey    string
	ModelName string
	Cli       *genai.Client
	Model     *genai.GenerativeModel

	enableContext   bool              // whether to use assistant context, default is false
	contextMessages []*ContextMessage // assistant context
}

// ContextMessage chat history message
type ContextMessage struct {
	Role    string `json:"role"` // "user" or "model"
	Content string `json:"content"`
}

// NewClient creates a new Google generative AI client.
func NewClient(apiKey string, opts ...ClientOption) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close closes the client.
func (c *Client) Close() error { _ = "STUB: not implemented"; return nil }

// Send sends a prompt to the gemini model and returns the response.
func (c *Client) Send(ctx context.Context, prompt string, files ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// SendStream sends a prompt to the gemini model and returns a channel of responses.
func (c *Client) SendStream(ctx context.Context, prompt string, files ...string) *aicli.StreamReply {
	_ = "STUB: not implemented"
	return nil
}

// ListModelNames lists the available models.
func (c *Client) ListModelNames(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListContextMessages list assistant context messages
func (c *Client) ListContextMessages() []*ContextMessage { _ = "STUB: not implemented"; return nil }

// RefreshContext refreshes assistant context
func (c *Client) RefreshContext() { _ = "STUB: not implemented"; return }

func (c *Client) appendAssistantContext(prompt string, replyContent string) {
	_ = "STUB: not implemented"
	return
}

func (c *Client) setMessages(prompt string, files ...string) ([]genai.Part, error) {
	_ = "STUB: not implemented"
	return nil,

		// history context
		nil
}

// file message

// user message
