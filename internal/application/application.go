/*
Package application run the application
*/
package application

import (
	"context"
	"fmt"

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
	filename, err := ValidateArgs(c.Args().Slice())
	if err != nil {
		return err
	}
	fmt.Printf("filename is %s\n", filename)
	return nil
}
func flags() []cli.Flag {
	return nil
}
