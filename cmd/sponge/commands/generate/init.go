package generate

import (
	"embed"
	"math/rand"
	"time"

	"github.com/go-dev-frame/sponge/pkg/gofile"
	"github.com/go-dev-frame/sponge/pkg/replacer"
)

const warnSymbol = "⚠ "

func init() {
	rand.Seed(time.Now().UnixNano()) //nolint
}

// Replacers replacer name
var Replacers = map[string]replacer.Replacer{}

// SpongeDir sponge directory
var SpongeDir = getHomeDir() + gofile.GetPathDelimiter() + ".sponge"

// Template information
type Template struct {
	Name     string
	FS       embed.FS
	FilePath string
}

// Init initializing the template
func Init() error {
	_ = "STUB: not implemented"
	// determine if the template file exists, if not, prompt to initialize first
	return nil
}

// InitFS initializing th FS templates
func InitFS(name string, filepath string, fs embed.FS) { _ = "STUB: not implemented"; return }

func isShowCommand() bool {
	_ = "STUB: not implemented"

	// sponge
	return false
}

// sponge init or sponge -h

func getHomeDir() string { _ = "STUB: not implemented"; return "" }
