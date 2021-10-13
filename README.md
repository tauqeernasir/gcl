# Go colorful logging

A dead simple go logging utility with or without colors.

### Usage

```go

package main

import (
	"os"

	"github.com/tauqeernasir/gcl/gcl"
)

var Logger *gcl.Logger

func init() {
	Logger = gcl.NewLogger(os.Stdout)
}

func main() {
	Logger.Info("I am an information message.")
	Logger.Error("I am an error message.")
	Logger.Success("I am a success message.")
	Logger.Warn("I am a warning message.")
	Logger.Fatal("I am a fatal message.")
}

```

Above code will generate the following output

<img width="716" alt="Screen Shot 2021-10-14 at 1 49 02 AM" src="https://user-images.githubusercontent.com/60596259/137217277-a4109709-a6ea-43bf-9c26-e8b0e00e53dd.png">

> Every method supports formatted versions as well, such as `Logger.Infof(format string, args ...interface{})` and `Logger.Errorf(format string, args ...interface{})`.

### Configurations

You can turn on/off some features of the logger by calling following methods on `Logger`
- WithColor() or WithoutColor()
- WithTimestamp() or WithoutTimestamp
- WithFileInfo() or WithoutFileInfo()

All methods return reference to the `Logger`, so they can easily be chained. Like `Logger.WithColor().WithoutTimestamp().WithoutFileInfo()`
