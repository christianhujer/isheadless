// +build !windows

package isheadless

import (
	"os"
	"testing"
)

func TestHeaded(t *testing.T) {
	os.Setenv("DISPLAY", ":0")
	if IsHeadless() {
		t.Error("Expected IsHeadless to return false because DISPLAY is set.")
	}
}

func TestHeadlessEmpty(t *testing.T) {
	os.Setenv("DISPLAY", "")
	if !IsHeadless() {
		t.Error("Expected IsHeadless to return true because DISPLAY is set but empty.")
	}
}

func TestHeadless(t *testing.T) {
	os.Unsetenv("DISPLAY")
	if !IsHeadless() {
		t.Error("Expected IsHeadless to return true because DISPLAY is unset.")
	}
}
