/* greet.go */
package main

import "os"
import "github.com/codegangsta/cli"

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"

	app.Name = "greet"
	app.Usage = "fight the loneliness!"

	app.Flags = []cli.Flag{
		cli.StringFlag{"lang", "english", "language for the greeting"},
	}
	app.Action = func(c *cli.Context) {
		name := "someone"
		if len(c.Args()) > 0 {
			name = c.Args()[0]
		}
		println(c.String("lang"))

		if c.String("lang") == "spanish" {
			println("Hola", name)
		} else {
			println("Hello", name)
		}
	}

	app.Commands = []cli.Command{
		{
			Name:      "add",
			ShortName: "a",
			Usage:     "add a task to the list",
			Action: func(c *cli.Context) {
				println("added task: ", c.Args()[0])
			},
		},
		{
			Name:      "complete",
			ShortName: "c",
			Usage:     "complete a task on the list",
			Action: func(c *cli.Context) {
				println("completed task: ", c.Args()[0])
			},
		},
	}

	app.Run(os.Args)
}
