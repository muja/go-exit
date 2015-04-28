package exit

import (
	"testing"
)

func TestRunHooksOnPanic(t *testing.T) {
	defer lastBreath()
	RunHooksOnPanic = false
	Hooks.Register(func() {
		t.Error("Hook was run on panic.")
	})
	panic("TestRunHooksOnPanic")
}

func TestPanicHook(t *testing.T) {
	called := false
	defer func() {
		if !called {
			t.Error("PanicHook was not called")
		}
	}()
	defer lastBreath()
	PanicHook = func(i interface{}) {
		called = true
	}
	panic("TestPanicHook")
}
