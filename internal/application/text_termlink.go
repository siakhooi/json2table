/*
Package application run the application
*/
package application

import (
	"os"

	"github.com/mattn/go-isatty"
	"github.com/savioxavier/termlink"
)

var stdoutIsTerminal = func() bool {
	fd := os.Stdout.Fd()
	return isatty.IsTerminal(fd) || isatty.IsCygwinTerminal(fd)
}

func applyLink(text, url string) string {
	if url == "" || !stdoutIsTerminal() {
		return text
	}
	return termlink.Link(text, url)
}
