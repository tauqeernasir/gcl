package main

import (
	"os"

	"github.com/tauqeernasir/gcl/gcl"
)

func main() {
	gcl := gcl.New(os.Stdout).WithTimestamp(true)
	gcl.Info("OK I am blue.")
	gcl.Info("OK I am Red.")
}
