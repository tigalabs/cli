package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	_ "github.com/github.com/tigateam/tiga-cli/commands/create"
	_ "github.com/github.com/tigateam/tiga-cli/commands/docker"
	_ "github.com/github.com/tigateam/tiga-cli/commands/info"
	_ "github.com/github.com/tigateam/tiga-cli/commands/run"
	_ "github.com/github.com/tigateam/tiga-cli/commands/ui"
	_ "github.com/github.com/tigateam/tiga-cli/commands/update"
)

func main() {
	// 帮助信息模板
	cli.AppHelpTemplate = fmt.Sprintf(`%s
AUTHOR : TigaTeam

GitHub: https://github.com/tigateam/tiga-cli
`, cli.AppHelpTemplate)

	app := &cli.App{
		Name:  "version",
		Usage: "show current binary version info | 显示当前版本信息",
		Action: func(c *cli.Context) error {
			fmt.Printf(`Tiga-CLI Tools v0.0.1
Install Path:E:\WorkSpace\GoPath\bin\tg.exe
Detail:	
  Go Version: go1.14
  TG version: tgv0.0.1
			`)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
