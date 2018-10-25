package main

import (
	. "fmt"
	. "github.com/christianhujer/isheadless"
	. "os"
)

func main() {
	headless := IsHeadless()
	Fprintf(Stderr, "%s: info: headless: %v\n", Args[0], headless)
	Exit(map[bool]int{true: 0, false: 1}[headless])
}
