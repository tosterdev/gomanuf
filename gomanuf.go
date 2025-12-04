package gomanuf

import (
	"embed"
	"regexp"
)

//go:embed manuf.txt
var manufFS embed.FS

var (
	data    map[string]string
	slash28 map[string]string
	slash36 map[string]string
)

var macRegex, _ = regexp.Compile("^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$")

func init() {
	loadManufData()
}
