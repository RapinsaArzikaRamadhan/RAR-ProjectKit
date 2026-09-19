package generator

import "embed"

//go:embed templates/go-sqlite/main.txt templates/go-sqlite/index.txt templates/go-sqlite/crud.txt templates/go-mysql/main.txt templates/go-mysql/index.txt templates/go-mysql/crud.txt
var Template embed.FS