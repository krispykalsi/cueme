//go:build windows
// +build windows

package main

import (
	"cueme/pkg/platform"
)

func main() {
	platform.EnableColorsInShell()
	runServer()
}
