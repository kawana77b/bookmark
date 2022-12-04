package main

import (
	"embed"

	cmd "github.com/kawana77b/bookmark/cmd"
)

var Version string = "v0.0.1"

//go:embed template
var tplFs embed.FS

func main() {
	cmd.Version = Version
	cmd.TplFs = tplFs

	cmd.Execute()
}
