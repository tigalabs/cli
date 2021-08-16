package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

//包初始化
func init() {
	rewriteHelpTemplate()
	rewriteVersion()
}

//重写帮助模板
func rewriteHelpTemplate() {
	cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}{{if .Version}}
VERSION:
   {{.Version}}
   {{end}}
USAGE:
   {{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
   {{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
   {{.Copyright}}
   {{end}}
`
}

//重写版本模板
func rewriteVersion() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s Version %s", c.App.Name, c.App.Version)
	}
}
