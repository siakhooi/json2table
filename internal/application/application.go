/*
Package application run the application
*/
package application

import (
	"context"
	"encoding/json"
	"fmt"
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
		ArgsUsage: "inputJsonFile",
		Version:   versioninfo.Version,
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

	// Check if file is readable
	_, err = os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot read file: %w", err)
	}

	// Read file contents
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	// Parse JSON
	var jsonData interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}

	// Pretty print JSON
	prettyJSON, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return fmt.Errorf("error formatting JSON: %w", err)
	}

	fmt.Println(string(prettyJSON))
	return nil
}
func flags() []cli.Flag {
	return nil
}
