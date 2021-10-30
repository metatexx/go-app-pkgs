package dbg

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"runtime"
	"strings"
)

const active = true

func Log(msg ...string) {
	if !active {
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
	txt := strings.TrimPrefix(strings.TrimPrefix(caller.Name(), "github.com/metatexx/gobro/"), "cmd/app/")
	if len(msg) > 0 {
		txt += ": " + strings.Join(msg, ", ")
	}
	app.Log(txt)
}

func Logf(fmt string, msg ...interface{}) {
	if !active {
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
	txt := strings.TrimPrefix(strings.TrimPrefix(caller.Name(), "github.com/metatexx/gobro/"), "cmd/app/")
	var args []interface{}
	args = append(args, txt)
	args = append(args, msg...)
	app.Logf("%s: "+fmt, args...)
}
