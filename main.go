package main

import (
	"os"

	"github.com/tauqeernasir/gcl/gcl"
)

func main() {
	l := gcl.NewLogger(os.Stdout).WithTimestamp().WithoutColor()

	l.WithFields(gcl.Fields{
		"name":  "tauqeer",
		"email": "tauqeer@g.com",
	}).Info("I am an information.")
	l.Infof("I am an information (%v=%v)", "key", "value")
	l.Warn("I am a warning")
	l.Warnf("I am an warning (%v=%v)", "key", "value")
	l.Success("I am an error.")
	l.Errorf("I am an error (%v=%v)", "key", "value")

	lc := gcl.NewLogger(os.Stdout).WithColor().WithPrettyJson()
	lc.Info("I am an information.")
	lc.Infof("I am an information (%v=%v)", "key", "value")
	lc.Warn("I am a warning")
	lc.WithFields(gcl.Fields{
		"name":  "tauqeer",
		"email": "tauqeer@g.com",
	}).Warnf("I am an warning (%v=%v)", "key", "value")
	lc.Success("I am an error.")
	lc.Errorf("I am an error (%v=%v)", "key", "value")
}
