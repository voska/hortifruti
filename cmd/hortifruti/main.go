package main

import (
	"github.com/voska/hortifruti"
	"github.com/voska/vtexkit/cli"
)

// version is injected at build time via -ldflags.
var version = "dev"

func main() {
	cli.Main(cli.App{
		Store:       hortifruti.Store,
		Version:     version,
		Description: "Hortifruti fresh produce CLI for humans and AI agents.",
	})
}
