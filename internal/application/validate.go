/*
Package application run the application
*/
package application

import "fmt"

// ValidateArgs of cli
func ValidateArgs(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("invalid arguments")
	}
	return nil
}
