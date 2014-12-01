package commitular

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func FormatSubject(changeType string, changeScope string, description string) string {
	return fmt.Sprintf("%s(%s): %s", changeType, changeScope, description)
}

func main() {
	app := cli.NewApp()
	app.Name = "commitular"
	app.Usage = "Make a git commit using the Angular commit message format"

	app.Commands = []cli.Command{
		{
			Name:  "feat",
			Usage: "new feature",
			Action: func(c *cli.Context) {
			},
		},
		{
			Name:  "fix",
			Usage: "bug fix",
			Action: func(c *cli.Context) {
			},
		},
		{
			Name:  "docs",
			Usage: "documentation only change",
			Action: func(c *cli.Context) {
			},
		},
		{
			Name:  "style",
			Usage: "style only change that does not affect the meaning of the code (white-space, formatting, missing punctuation, etc)",
			Action: func(c *cli.Context) {
			},
		},
		{
			Name:  "refactor",
			Usage: "change that neither fixes a bug or adds a feature",
			Action: func(c *cli.Context) {
			},
		},
		{
			Name:  "perf",
			Usage: "change that improves performance",
			Action: func(c *cli.Context) {
			},
		},
		{
			Name:  "test",
			Usage: "change that improves test coverage",
			Action: func(c *cli.Context) {
			},
		},
		{
			Name:  "chore",
			Usage: "change the affects build processes, tooling, or libraries",
			Action: func(c *cli.Context) {
			},
		},
	}

	app.Run(os.Args)
}
