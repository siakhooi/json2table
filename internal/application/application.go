/*
Package application run the application
*/
package application

import (
	"context"

	"github.com/siakhooi/json2table/internal/version"
	"github.com/urfave/cli/v3"
)

/*
Run start application
*/
func Run(args []string) error {
	cmd := &cli.Command{
		Name:      "json2table",
		Usage:     "convert json data to tabular format",
		ArgsUsage: "inputJsonFile",
		Version:   version.Version(),
		Action:    action,
		Flags:     flags(),
	}
	return cmd.Run(context.Background(), args)

}

func action(_ context.Context, c *cli.Command) error {
	args := c.Args().Slice()
	if len(args) == 0 {
		return cli.ShowAppHelp(c)
	}

	return ValidateArgs(args)
}
func flags() []cli.Flag {
	return nil
}
