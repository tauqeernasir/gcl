package main

import (
	"os"

	"github.com/tauqeernasir/gcl/gcl"
)

func main() {
	gcl := gcl.New(os.Stdout).WithTimestamp()
	gcl.Info("OK I am blue.")
	gcl.WithoutColor()
	gcl.Info("OK I am Red.")
}
