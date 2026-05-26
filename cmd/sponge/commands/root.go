// Package commands are subcommands of the sponge command.
package commands

import (
	"github.com/spf13/cobra"
)

var (
	version     = "v0.0.0"
	versionFile = GetSpongeDir() + "/.sponge/.github/version"
)

// NewRootCMD command entry
func NewRootCMD() *cobra.Command { _ = "STUB: not implemented"; return nil }

func getVersion() string { _ = "STUB: not implemented"; return "" }

// GetSpongeDir get sponge home directory
func GetSpongeDir() string { _ = "STUB: not implemented"; return "" }
