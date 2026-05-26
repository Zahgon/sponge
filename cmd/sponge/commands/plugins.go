package commands

import (
	"github.com/spf13/cobra"
)

var pluginNames = []string{
	"go",
	"protoc",
	"protoc-gen-go",
	"protoc-gen-go-grpc",
	"protoc-gen-validate",
	"protoc-gen-gotag",
	"protoc-gen-go-gin",
	"protoc-gen-go-rpc-tmpl",
	"protoc-gen-json-field",
	"protoc-gen-openapiv2",
	"protoc-gen-doc",
	"swag",
	//"golangci-lint",
	//"go-callvis",
}

var installPluginCommands = map[string]string{
	"go":                     "go: please install manually yourself, download url is https://go.dev/dl/ or https://golang.google.cn/dl/",
	"protoc":                 "protoc: please install manually yourself, download url is https://github.com/protocolbuffers/protobuf/releases/tag/v31.1",
	"protoc-gen-go":          "google.golang.org/protobuf/cmd/protoc-gen-go@latest",
	"protoc-gen-go-grpc":     "google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest",
	"protoc-gen-validate":    "github.com/envoyproxy/protoc-gen-validate@latest",
	"protoc-gen-gotag":       "github.com/srikrsna/protoc-gen-gotag@latest",
	"protoc-gen-go-gin":      "github.com/go-dev-frame/sponge/cmd/protoc-gen-go-gin@latest",
	"protoc-gen-go-rpc-tmpl": "github.com/go-dev-frame/sponge/cmd/protoc-gen-go-rpc-tmpl@latest",
	"protoc-gen-json-field":  "github.com/go-dev-frame/sponge/cmd/protoc-gen-json-field@latest",
	"protoc-gen-openapiv2":   "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest",
	"protoc-gen-doc":         "github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest",
	"swag":                   "github.com/swaggo/swag/cmd/swag@v1.8.12",
	//"golangci-lint":          "github.com/golangci/golangci-lint/cmd/golangci-lint@latest",
	//"go-callvis":             "github.com/ofabry/go-callvis@latest",
}

const (
	installedSymbol = "✔ "
	lackSymbol      = "❌ "
	warnSymbol      = "⚠ "
)

// PluginsCommand plugins management
func PluginsCommand() *cobra.Command { _ = "STUB: not implemented"; return nil }

func checkInstallPlugins() (installedNames []string, lackNames []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

func showDependencyPlugins(installedNames []string, lackNames []string) {
	_ = "STUB: not implemented"
	return
}

func installPlugins(lackNames []string) { _ = "STUB: not implemented"; return }

//nolint

func adaptInternalCommand(name string, pkgAddr string) string { _ = "STUB: not implemented"; return "" }

func filterLackNames(lackNames []string, skipPluginName string) []string {
	_ = "STUB: not implemented"
	return nil
}
