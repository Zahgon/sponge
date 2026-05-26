package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := perftestCommand()
	if err := cmd.Execute(); err != nil {
		cmd.PrintErrln("Error:", err)
		os.Exit(1)
	}
}

func perftestCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }
