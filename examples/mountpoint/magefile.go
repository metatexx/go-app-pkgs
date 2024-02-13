//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

var Default = Run

func Run() error {
	err := sh.RunWithV(map[string]string{"GOOS": "js", "GOARCH": "wasm"}, "go", "build", "-o", "web/app.wasm", "main.go")
	if err != nil {
		return err
	}
	return sh.RunV("go", "run", "main.go")
}
