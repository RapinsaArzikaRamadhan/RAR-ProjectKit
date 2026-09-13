package generator

import "embed"

//go:embed templates/go-sqlite/main.txt templates/go-sqlite/index.txt templates/go-sqlite/crud.txt
var Template embed.FS