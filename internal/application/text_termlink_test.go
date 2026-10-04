package application

import (
	"os"
	"strings"
	"testing"

	"github.com/savioxavier/termlink"
)

func TestGetLink_EmptyURLReturnsText(t *testing.T) {
	text := "hello"
	got := applyLink(text, "")
	if got != text {
		t.Fatalf("GetLink(%q, %q) = %q, want %q", text, "", got, text)
	}
}

func TestGetLink_WithURLReturnsTerminalLink(t *testing.T) {
	original := stdoutIsTerminal
	stdoutIsTerminal = func() bool { return true }
	t.Cleanup(func() { stdoutIsTerminal = original })

	text := "repo"
	url := "https://example.com"
	want := termlink.Link(text, url)

	got := applyLink(text, url)
	if got != want {
		t.Fatalf("GetLink(%q, %q) = %q, want %q", text, url, got, want)
	}
}

func TestGetLink_SuppressesHyperlinkWhenStdoutIsNotTerminal(t *testing.T) {
	originalStdout := os.Stdout
	readPipe, writePipe, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stdout pipe: %v", err)
	}
	os.Stdout = writePipe
	t.Cleanup(func() {
		os.Stdout = originalStdout
		_ = readPipe.Close()
		_ = writePipe.Close()
	})

	text := "repo"
	url := "https://example.com"
	got := applyLink(text, url)
	if got != text {
		t.Fatalf("GetLink(%q, %q) = %q, want plain %q", text, url, got, text)
	}
	if strings.Contains(got, "\x1b") || strings.Contains(got, url) {
		t.Fatalf("GetLink(%q, %q) = %q, want no hyperlink escape codes", text, url, got)
	}
}
