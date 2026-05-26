package benchmark

const (
	packagePattern = `\npackage (.*);`
	servicePattern = `\nservice (\w+)`
	methodPattern  = `rpc (\w+)`
)

func getName(data []byte, pattern string) string { _ = "STUB: not implemented"; return "" }

func getMethodNames(data []byte, methodPattern string) []string {
	_ = "STUB: not implemented"
	return nil
}

// match name, not case-sensitive
func matchName(names []string, name string) string { _ = "STUB: not implemented"; return "" }
