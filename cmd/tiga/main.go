package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	app = &cli.App{
		Name:    "Tiga CLI",
		Version: "v0.0.8",
		Usage:   "Tiga Ecosystem Command Line Tool",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Misitebao",
				Email: "i@misitebao.com",
			},
		},
		Copyright: "https://github.com/tigateam",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "language",
				Aliases:     []string{"lang"},
				Value:       "en",
				DefaultText: "en",
				Usage:       "Specify the display `language`",
				//Destination: &language,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Print(c.String("language"))
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "config",
				Aliases: []string{"conf"},
				Usage:   "Configuration items of the Tiga CLI Tool",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "username",
						Aliases:     []string{},
						Value:       "",
						DefaultText: "",
						Usage:       "Setup username",
						//Destination: &language,
					},
					&cli.StringFlag{
						Name:        "email",
						Aliases:     []string{},
						Value:       "",
						DefaultText: "",
						Usage:       "Setup email address",
						//Destination: &language,
					},
				},
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "create",
				Aliases: []string{},
				Usage:   "Create a project",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "name",
						Aliases:     []string{"n"},
						Value:       "tiga-app",
						DefaultText: "tiga-app",
						Usage:       "Specify project `name`",
						//Destination: &language,
					},
					&cli.StringFlag{
						Name:        "template",
						Aliases:     []string{"tmpl", "t"},
						Value:       "default",
						DefaultText: "default",
						Usage:       "Specify a new project templateName",
						//Destination: &language,
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Print(c.String("name"))
					fmt.Print(c.String("template"))

					return nil
				},
			},
			{
				Name:    "run",
				Aliases: []string{},
				Usage:   "Run the project",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "build",
				Aliases: []string{},
				Usage:   "Build the project",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "update",
				Aliases: []string{"u", "upgrade"},
				Usage:   "Upgrade to the latest version",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	}
)

//程序入口
func main() {
	//运行程序
	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
