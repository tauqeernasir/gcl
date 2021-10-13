package gcl

import (
	"fmt"
	"time"
)

func getCurrentTime(now time.Time) []byte {
	year, month, day := now.Date()
	hours, minutes, seconds := now.Clock()

	return []byte(fmt.Sprintf("[%v/%v/%v %v:%v:%v]", year, int(month), day, hours, minutes, seconds))
}
