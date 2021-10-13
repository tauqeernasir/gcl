package main

import (
	"os"

	"github.com/tauqeernasir/gcl/gcl"
)

func main() {
	gcl := gcl.New(os.Stdout).WithTimestamp()
	gcl.Info("I am an information.")
	gcl.Warn("I am a warning")
	gcl.Error("I am an error.")
}
