//go:build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/sh"
	"os"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
//var Default = Run

func Build() error {
	err := sh.RunV("go", "build", "-o", "./mounpoint", "main.go")
	if err != nil {
		return err
	}
	return sh.RunWithV(map[string]string{"GOOS": "js", "GOARCH": "wasm"}, "go", "build", "-o", "web/app.wasm", "main.go")
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	_ = os.RemoveAll("./web")
	_ = os.Remove("./mountpoint")
}
