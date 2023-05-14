package database

import "embed"

//go:embed *.sql
var Models embed.FS
