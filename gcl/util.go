package gcl

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func getCurrentTime(now time.Time) []byte {
	year, month, day := now.Date()
	hours, minutes, seconds := now.Clock()

	return []byte(fmt.Sprintf("[%v/%v/%v %v:%v:%v]", year, int(month), day, hours, minutes, seconds))
}

func getFileInfo() (
	file string,
	line int,
	fn string,
) {
	var (
		pc uintptr
		ok bool
	)
	if pc, file, line, ok = runtime.Caller(3); !ok {
		file = "unknown file"
		fn = "unknown function"
		line = 0
	} else {
		file = filepath.Base(file)
		arrFn := strings.Split(runtime.FuncForPC(pc).Name(), "/")
		fn = arrFn[len(arrFn)-1]
	}
	return
}
