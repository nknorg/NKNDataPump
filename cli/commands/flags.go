package commands

import "github.com/urfave/cli"

func SetFlags(app *cli.App) {
	app.Flags = []cli.Flag{
		new(SetPortFlag).NewFlag(),
		new(SetConfigFlag).NewFlag(),
	}
}
