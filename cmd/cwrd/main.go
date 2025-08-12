package main

import (
	"cwrd/internal/cmd"
	"fmt"
	"os"

	"github.com/alecthomas/kong"
)

const appName = "HAM CW Random Text Generator"
const appDesc = "Used for HAM amateurs to create random characters"

var version = "v0.0.1"

type VersionFlag string

func (v VersionFlag) Decode(_ *kong.DecodeContext) error { return nil }

func (v VersionFlag) IsBool() bool { return true }

func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

type CLI struct {
	cmd.Globals
	Rand    cmd.RandCommand `cmd:"" default:"withargs" help:"Used to create."`
	Version VersionFlag     `help:"Print version information and quit" short:"v" name:"version"`
}

func run() error {
	cli := CLI{
		Version: VersionFlag(version),
	}

	ctx := kong.Parse(&cli,
		kong.Name(appName),
		kong.Description(appDesc),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.DefaultEnvars(appName),
		kong.Vars{
			"version": string(cli.Version),
		})

	err := ctx.Run(&cli.Globals)
	ctx.FatalIfErrorf(err)
	return nil
}

func main() {
	if err := run(); err != nil {
		os.Exit(-1)
	}
}
