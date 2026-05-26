package gemini

const (
	Model25Flash = "gemini-2.5-flash"
	Model25Pro   = "gemini-2.5-pro"

	DefaultModel = Model25Flash

	RoleUser  = "user"
	RoleModel = "model"
)

// ClientOption is a function that sets a Client option.
type ClientOption func(*Client)

func defaultClientOptions() *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) apply(opts ...ClientOption) { _ = "STUB: not implemented"; return }

// WithModel sets the model name
func WithModel(name string) ClientOption { _ = "STUB: not implemented"; return *new(ClientOption) }

// WithEnableContext enable assistant context
func WithEnableContext() ClientOption { _ = "STUB: not implemented"; return *new(ClientOption) }

// WithInitialContextMessages sets assistant initial context messages
func WithInitialContextMessages(messages ...*ContextMessage) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

// default role is user
