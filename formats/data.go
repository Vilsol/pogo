package formats

import (
	"embed"
	"io/fs"
)

//go:embed formats/*.json
var rawEmbeddedFormats embed.FS

var EmbeddedFormats fs.FS

func init() {
	// we embedded formats/*.json, but we actually just want *.json
	subfs, err := fs.Sub(rawEmbeddedFormats, "formats")
	if err != nil {
		panic(err)
	}
	EmbeddedFormats = subfs
}
