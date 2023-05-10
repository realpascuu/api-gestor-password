package certs

import "embed"

//go:embed *.crt *.key
var Certs embed.FS
