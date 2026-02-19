/*
Package application run the application
*/
package application

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/siakhooi/json2table/internal/versioninfo"
	"github.com/urfave/cli/v3"
)

/*
Run start application
*/
func Run(args []string) error {
	cmd := &cli.Command{
		Name:      "json2table",
		Usage:     "convert json data to tabular format",
		ArgsUsage: "[dataFile]",
		Version:   versioninfo.Version,
		Action:    action,
		Flags:     flags(),
	}
	return cmd.Run(context.Background(), args)

}

func action(_ context.Context, c *cli.Command) error {
	if c.Bool("build") {
		versioninfo.PrintBuildInfo()
		return nil
	}

	args, err := ParseArguments(c)
	if err != nil {
		return err
	}
	// Read spec data
	specData, err := ReadSpec(args.SpecFile)
	if err != nil {
		return err
	}
	// Parse and validate the spec
	spec, err := ParseAndValidateSpec(specData)
	if err != nil {
		return err
	}

	// Pretty print spec JSON
	prettyJSON, err := json.MarshalIndent(spec, "", "  ")
	if err != nil {
		return fmt.Errorf("error formatting spec file: %w", err)
	}

	fmt.Println(string(prettyJSON))

	data, err := ReadData(args.DataFile)
	if err != nil {
		return err
	}

	// Parse JSON
	var jsonData interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}

	// Pretty print JSON
	prettyJSON1, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return fmt.Errorf("error formatting JSON: %w", err)
	}

	fmt.Println(string(prettyJSON1))
	return nil
}
func flags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  "build",
			Usage: "print build info and exit",
		},
		&cli.StringFlag{
			Name:    "spec",
			Aliases: []string{"s"},
			Usage:   "read spec from specFile.json, or from environment variable JSON2TABLE_SPEC_FILE if not provided",
		},
	}
}
