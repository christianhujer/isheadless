// Package isheadless provides a function to find out whether the environment of the process is headless.
package isheadless

import (
	"os"
)

// IsHeadless returns whether or not the environment of the current process is headless.
// That is, whether or not a display, keyboard, and mouse can be supported in this environment.
func IsHeadless() bool {
	_, headed := os.LookupEnv("DISPLAY")
	return !headed
}
