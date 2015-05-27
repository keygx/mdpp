package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "mdpp"
	app.Version = Version
	app.Usage = "Search .mobileprovision file from '~/Library/MobileDevice/Provisioning Profiles/' directory"
	app.Author = "keygx"
	app.Email = "https://twitter.com/keygx"
	app.Commands = Commands

	app.Run(os.Args)
}
