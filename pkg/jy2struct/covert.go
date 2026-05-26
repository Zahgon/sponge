// Package jy2struct is a library for generating go struct code, supporting json and yaml.
package jy2struct

// Args  convert arguments
type Args struct {
	Format    string // document format, json or yaml
	Data      string // json or yaml content
	InputFile string // file
	Name      string // name of structure
	SubStruct bool   // are sub-structures separated
	Tags      string // add additional tags, multiple tags separated by commas

	tags          []string //nolint
	convertFloats bool
	parser        Parser
}

func (j *Args) checkValid() error { _ = "STUB: not implemented"; return nil }

// Convert json or yaml to go struct
func Convert(args *Args) (string, error) { _ = "STUB: not implemented"; return "", nil }
