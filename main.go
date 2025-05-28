package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/lomifile/create-project/project"
	"github.com/urfave/cli/v3"
)

func main() {
	home, _ := os.UserHomeDir()
	defaultPath := fmt.Sprintf("%s/%s", home, "Projects")

	cmd := &cli.Command{
		EnableShellCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "js",
				Usage: "Specify project type",
			},
			&cli.StringFlag{
				Name:    "path",
				Aliases: []string{"p"},
				Value:   "",
				Usage:   "Specify path",
			},
		},
		Action: func(c context.Context, cmd *cli.Command) error {
			name := "sample-test"

			if cmd.NArg() > 0 {
				name = cmd.Args().Get(0)
			}

			lang := cmd.String("lang")
			framework := cmd.String("framework")
			if (lang == project.TS || lang == project.JS) && len(framework) > 0 {
				log.Printf("Specified framework is %s", framework)
			}

			project := &project.Project{
				Name:      name,
				Type:      lang,
				Framework: framework,
				Path:      defaultPath,
				Deps: map[string][]string{
					"dev":    {},
					"normal": {},
				},
			}

			project.Create()
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
