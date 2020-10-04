// package tg

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/urfave/cli/v2"
// )

// func main() {
// 	cli.AppHelpTemplate = fmt.Sprintf(`%s
// AUTHOR : TigaTeam

// WEBSITE: https://github.com/tigateam/tiga-cli
// `, cli.AppHelpTemplate)

// 	app := &cli.App{
// 		Name:  "version",
// 		Usage: "show current binary version info | 显示当前版本信息",
// 		Action: func(c *cli.Context) error {
// 			fmt.Println("v0.0.1")
// 			return nil
// 		},
// 	}

// 	err := app.Run(os.Args)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
