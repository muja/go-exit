package exit

import "os"

type hooks []func()

// Register hooks that will be run before the
// client terminates. These will only be run
// if RunHooks == true (by default) when no panic occurs
// or RunHooksOnPanic == true (by default) when panic occurrs.
func (h *hooks) Register(f func()) {
	*h = append(*h, f)
}

// Whether the exit should recover from panics
var Recover = true

// The status code the client will terminate with
// (when no panic occurred.) Can be set any time.
var Status = 0

// The status code the client will terminate with
// when a panic has occurred. Can be set any time.
var StatusOnPanic = 2

// Hooks to be run before exiting.
var Hooks = make(hooks, 0)

// Whether hooks should be run when no panic has occurred.
// Can be set any time.
var RunHooks = true

// Whether hooks should be run on panic. Can be set any time.
var RunHooksOnPanic = true

// A hook that will be run if a panic has occurred.
// The value passed to panic will be passed here.
var PanicHook = func(i interface{}) {}

func lastBreath() {
	if Recover {
		if r := recover(); r != nil {
			RunHooks = RunHooksOnPanic
			Status = StatusOnPanic
			PanicHook(r)
		}
	}
	if RunHooks {
		for _, hook := range Hooks {
			hook()
		}
	}
}

func Exit() {
	lastBreath()
	os.Exit(Status)
}
