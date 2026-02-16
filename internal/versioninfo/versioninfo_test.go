package versioninfo

import "testing"

const expectedjson2tableVersion = "0.0.1"

func TestGetVersion(t *testing.T) {
	actual := Version()
	expected := expectedjson2tableVersion

	if actual != expected {
		t.Errorf("GetVersion() = %q, want %q", actual, expected)
	}
}
