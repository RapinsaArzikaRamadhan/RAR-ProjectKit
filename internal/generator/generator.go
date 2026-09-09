package generator

import "embed"

//go:embed templates/go/main.txt templates/go/index.txt templates/go/crud.txt
var Template embed.FS