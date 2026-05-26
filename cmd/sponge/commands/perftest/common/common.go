package common

const (
	BodyTypeJSON = "application/json"
	BodyTypeForm = "application/x-www-form-urlencoded"
	BodyTypeText = "text/plain"
)

var CommandPrefix = "sponge perftest"

// SetCommandPrefix sets the command prefix for the perftest command.
func SetCommandPrefix(name string) { _ = "STUB: not implemented"; return }

// nolint
func ParseHTTPParams(method string, headers []string, body string, bodyFile string) ([]byte, map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func trimString(s string) string { _ = "STUB: not implemented"; return "" }

// CheckBodyParam checks if the body parameter is provided in JSON or file format.
func CheckBodyParam(bodyJSON string, bodyFile string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func isValidJSON(data []byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NewID generates a new ID for each request.
func NewID() int64 { _ = "STUB: not implemented"; return 0 }

// NewStringID Generate a string ID, the hexadecimal form of NewID(), total 16 bytes.
func NewStringID() string { _ = "STUB: not implemented"; return "" }

// CheckPortInUse checks if the given port is in use, if not, it returns a new available port.
func CheckPortInUse(port string) string { _ = "STUB: not implemented"; return "" }
