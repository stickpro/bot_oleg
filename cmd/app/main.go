package main

import (
	"bot-oleg/cmd/console"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	application := &cli.App{
		Name:                 "Bot Oleg",
		Description:          "This is an API for DV Merchant",
		Suggest:              true,
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			cli.HelpFlag,
			cli.VersionFlag,
			cli.BashCompletionFlag,
		},
		Commands: console.InitCommands(),
	}
	if err := application.Run(os.Args); err != nil {
		_, _ = fmt.Println(err.Error())
		os.Exit(1)
	}
}
