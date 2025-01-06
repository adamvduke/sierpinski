package static

import (
	"embed"
	"io/fs"
)

//go:embed public
var public embed.FS

// PublicFS returns an fs.FS that has a root of the application's public directory.
func PublicFS() (fs.FS, error) {
	return fs.Sub(public, "public")
}
