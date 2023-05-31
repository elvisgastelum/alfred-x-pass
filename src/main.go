package main

import (
	"os"
	"time"

	filter "github.com/elvisgastelum/pasawutil/src/filter"
	glf "github.com/elvisgastelum/pasawutil/src/get-login-field"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "pasawutil",
		Version:   "v0.0.1",
		Copyright: "(c) 2023 Elvis Gastelum",
		Usage:     "Pass Store Alfred Workflow Utiliy Binary",
		Compiled:  time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Elvis Gastelum",
				Email: "elvisgastelum@outlook.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:      "filter",
				Aliases:   []string{"f"},
				Usage:     "pasawutil [filter|f] [query]",
				UsageText: "pass a query to filter the password store",
				Action: func(cCtx *cli.Context) error {
					filter.RunFilter()
					return nil
				},
			},
			{
				Name:      "get-login",
				Aliases:   []string{"l"},
				Usage:     "pasawutil [get-login|l] [query]",
				UsageText: "pass a query to filter the password store and get the login field",
				Action: func(cCtx *cli.Context) error {
					glf.GetLoginField(
						cCtx.Args().Get(0),
					)
					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}
