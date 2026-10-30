package main

import (
	// Modules of packwiz
	"github.com/kek-team/packwiz/cmd"
	_ "github.com/kek-team/packwiz/curseforge"
	_ "github.com/kek-team/packwiz/github"
	_ "github.com/kek-team/packwiz/migrate"
	_ "github.com/kek-team/packwiz/modrinth"
	_ "github.com/kek-team/packwiz/settings"
	_ "github.com/kek-team/packwiz/url"
	_ "github.com/kek-team/packwiz/utils"
)

func main() {
	cmd.Execute()
}
