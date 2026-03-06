/*
Package application run the application
*/
package application

import (
	"github.com/savioxavier/termlink"
)

// GetLink returns a terminal link if the url is not empty, otherwise it returns the text as is
func GetLink(text string, url string) string {
	if url == "" {
		return text
	}
	return termlink.Link(text, url)
}
