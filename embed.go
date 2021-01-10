package embedexample

import (
	"embed"
)

//go:embed static
var StaticAsset embed.FS

//go:embed template
var TemplateFS embed.FS
