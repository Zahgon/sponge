// Package conf is parsing yaml, json, toml configuration files to go struct.
package conf

// Parse configuration files to struct, including yaml, toml, json, etc., and turn on listening for configuration file changes if fs is not empty
func Parse(configFile string, obj interface{}, reloads ...func()) error {
	_ = "STUB: not implemented"
	return nil
}

// excluding suffix names

// path
// file name
// get the configuration type from the file name

// ParseConfigData parse data to struct, parameter format is the configuration file format, such as "yaml", "json", "toml"
func ParseConfigData(data []byte, format string, obj interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// listening for profile updates
func watchConfig(obj interface{}, reloads ...func()) {
	_ = "STUB: not implemented"

	// Note: OnConfigChange is called twice on Windows
	return
}

// reset object

// Show print configuration information (hide sensitive fields)
func Show(obj interface{}, fields ...string) string { _ = "STUB: not implemented"; return "" }

func hideSensitiveFields(line string, fields ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// replace dsn

// replace dsn password
func replaceDSN(str string) string { _ = "STUB: not implemented"; return "" }
