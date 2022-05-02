package dbg

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"runtime"
	"strings"
)

// Enabled is used to anable or disable all logging
var Enabled = true

// TrimPrefixes defines the prefixes to strip from the output
// so it is less verbose
var TrimPrefixes = []string{""}

// Log logs the string(s) together with the caller information
func Log(msg ...interface{}) {
	if !Enabled {
		return
	}
	fpcs := make([]uintptr, 1)

	// Skip 2 levels to get the caller
	n := runtime.Callers(2, fpcs)
	if n == 0 {
		app.Log("MSG: NO CALLER")
	}

	caller := runtime.FuncForPC(fpcs[0] - 1)
	if caller == nil {
		app.Log("MSG CALLER WAS NIL")
	}
	// Print the name of the function (but trim to our package names)
	txt := caller.Name()
	for _, prefix := range TrimPrefixes {
		txt = strings.TrimPrefix(txt, prefix)
	}
	if len(msg) > 0 {
		txt += ":"
		var args []interface{}
		args = append(args, txt)
		args = append(args, msg...)
		app.Log(args...)
	} else {
		app.Log(txt)
	}
}

// Logf is like Log() but with a sprintf style format string
func Logf(fmt string, msg ...interface{}) {
	if !Enabled {
		return
	}
	fpcs := make([]uintptr, 1)

	// Skip 2 levels to get the caller
	n := runtime.Callers(2, fpcs)
	if n == 0 {
		app.Log("MSG: NO CALLER")
	}

	caller := runtime.FuncForPC(fpcs[0] - 1)
	if caller == nil {
		app.Log("MSG CALLER WAS NIL")
	}
	// Print the name of the function (but trim to our package names)
	txt := caller.Name()
	for _, prefix := range TrimPrefixes {
		txt = strings.TrimPrefix(txt, prefix)
	}
	var args []interface{}
	args = append(args, txt)
	args = append(args, msg...)
	app.Logf("%s: "+fmt, args...)
}

// Logc uses console.log() and makes it easy to debug jsValue in the browser
func Logc(msg interface{}) {
	if !Enabled {
		return
	}
	app.Window().Get("console").Call("log", msg)
}
