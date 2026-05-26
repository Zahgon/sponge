package commands

import (
	"github.com/spf13/cobra"
)

// UpgradeCommand upgrade sponge binaries
func UpgradeCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func runUpgrade(targetVersion string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func runUpgradeCommand(targetVersion string) error { _ = "STUB: not implemented"; return nil }

//nolint

// copy the template files to a temporary directory
func copyToTempDir(targetVersion string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// use the first $GOPATH

// find the new version of the sponge code directory

func executeCommand(name string, args ...string) error { _ = "STUB: not implemented"; return nil }

//nolint

func adaptPathDelimiter(filePath string) string { _ = "STUB: not implemented"; return "" }

func getLatestVersion(s string) string { _ = "STUB: not implemented"; return "" }

func updateSpongeInternalPlugin(targetVersion string) error { _ = "STUB: not implemented"; return nil }

//nolint

//nolint

// v1.x.x version does not support protoc-gen-json-field

//nolint

// v1 >= v2 return true
// v1 < v2 return false
func compareVersion(v1, v2 string) bool { _ = "STUB: not implemented"; return false }
