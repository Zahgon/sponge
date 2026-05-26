// Package template provides commands to generate custom code.
package template

import (
	"regexp"
	"strings"
)

func parseFields(jsonFile string) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mergeFields(m1 map[string]interface{}, m2 map[string]interface{}) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func listTemplateFiles(builder *strings.Builder, files []string) { _ = "STUB: not implemented"; return }

func listFields(builder *strings.Builder, fields map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

var regPackage = regexp.MustCompile(`(?m)^package\s+([a-zA-Z0-9._]+);`)

func copyProtoFileToDir(protoFile string, targetDir string) error {
	_ = "STUB: not implemented"
	return nil
}
