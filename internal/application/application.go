/*
Package application run the application
*/
package application

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

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
	// If --build was provided, print build info and exit
	if c.Bool("build") {
		versioninfo.PrintBuildInfo()
		return nil
	}

	// Read and validate spec (mandatory)
	specFile := c.String("spec")
	spec, err := ReadSpec(specFile)
	if err != nil {
		return err
	}
	// Pretty print spec JSON
	prettyJSON, err := json.MarshalIndent(spec, "", "  ")
	if err != nil {
		return fmt.Errorf("error formatting spec file: %w", err)
	}

	fmt.Println(string(prettyJSON))

	filename, err := ValidateArgs(c.Args().Slice())
	if err != nil {
		return err
	}

	var data []byte
	if filename == "-" {
		// read from stdin
		data, err = io.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("error reading stdin: %w", err)
		}
	} else {
		// Check if file is readable
		_, err = os.Open(filename)
		if err != nil {
			return fmt.Errorf("cannot read file: %w", err)
		}

		// Read file contents
		data, err = os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}
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
