package main

import (
	"os"

	"github.com/tauqeernasir/gcl/gcl"
)

func main() {
	gcl := gcl.New(os.Stdout).WithTimestamp().WithColor()
	gcl.Info("I am an information.")
	gcl.Infof("I am an information (%v=%v)", "key", "value")
	gcl.Warn("I am a warning")
	gcl.Warnf("I am an warning (%v=%v)", "key", "value")
	gcl.Success("I am an error.")
	gcl.Errorf("I am an error (%v=%v)", "key", "value")
}
