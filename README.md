Exit - Configure your app's termination behaviour.
==================================================

# Table of contents

1. Introduction
2. Usage
3. Contributing
4. Reporting bugs

-------------------

# 1. Introduction

This package provides a way of controlling your exit behaviour.  
It's often important to let some goroutines finish before main returns,
or you want to recover from a panic and exit gracefully.
There are numerous things that you might want to do, without having to
cluster your main function with all this logic.

# 2. Usage

```go
package main

import (
  "os"
  "github.com/muja/go-exit"
  "log"
)

func main() {
  // Make sure we exit via the configurable exit module.
  defer exit.Exit()

  // Enable recovering from panics:
  exit.Recover = true

  // You might register a panic hook that gets called in the Exit()
  // function when the application panics, for example, the default hook:
  exit.PanicHook = func(i interface{}) {
    log.Printf("Fatal error: %v\n", i)
  }

  // Set the status code (by default this is 0)
  exit.Status = 2

  // Register hooks that will be run before termination:
  exit.Hooks.Register(func() {
    // Some crucial filesystem cleanup / closing / etc.
  })

  // Disable hooks to be run (true by default):
  exit.RunHooks = false

  // Disable hooks to be run on panic (true by default):
  exit.RunHooksOnPanic = false
}
```

# 3. Contributing

Contributions are welcome! Fork -> Patch -> Push -> Pull request.

# 4. Bug report / suggestions

Just create an issue! I will try to reply as soon as possible.
