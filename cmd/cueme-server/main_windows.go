//go:build windows

package main

import (
	"cueme/pkg/platform"
)

func init() {
	platform.EnableColorsInShell()
}
