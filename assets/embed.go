package assets

import "embed"

var (
	//go:embed textures
	textures embed.FS

	//go:embed shaders
	shaders embed.FS
)
