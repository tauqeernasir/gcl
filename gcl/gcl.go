package gcl

import (
	"fmt"
	"io"
	"sync"
	"time"
)

var infoPrefixPlain = []byte("[INFO] ")
var warnPrefixPlain = []byte("[WARN] ")
var errorPrefixPlain = []byte("[WARN] ")

type Prefix struct {
	Plain   []byte
	Colored []byte
}

var (
	InfoPrefix = Prefix{
		Plain:   infoPrefixPlain,
		Colored: Blue(infoPrefixPlain),
	}

	WarnPrefix = Prefix{
		Plain:   warnPrefixPlain,
		Colored: Yellow(warnPrefixPlain),
	}

	ErrorPrefix = Prefix{
		Plain:   errorPrefixPlain,
		Colored: Red(errorPrefixPlain),
	}
)

type Logger struct {
	mu        sync.RWMutex
	color     bool
	timestamp bool
	buf       Buffer
	out       io.Writer
}

func (l *Logger) Info(text string) {
	l.Log(InfoPrefix, text)
}

func (l *Logger) Infof(text string, args ...interface{}) {
	l.Log(InfoPrefix, fmt.Sprintf(text, args...))
}

func (l *Logger) Warn(text string) {
	l.Log(WarnPrefix, text)
}

func (l *Logger) Warnf(text string, args ...interface{}) {
	l.Log(WarnPrefix, fmt.Sprintf(text, args...))
}

func (l *Logger) Error(text string) {
	l.Log(ErrorPrefix, text)
}

func (l *Logger) Errorf(text string, args ...interface{}) {
	l.Log(ErrorPrefix, fmt.Sprintf(text, args...))
}

func (l *Logger) WithTimestamp() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.timestamp = true
	return l
}

func (l *Logger) WithoutTimestamp() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.timestamp = false
	return l
}

func (l *Logger) WithColor() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.color = true
	return l
}

func (l *Logger) WithoutColor() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.color = false
	return l
}

func (l *Logger) Log(prefix Prefix, text string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()

	l.buf.Reset()

	if l.color {
		l.buf.Append(prefix.Colored)
	} else {
		l.buf.Append(prefix.Plain)
	}

	if l.timestamp {
		if l.color {
			l.buf.Append(ColorPurple)
		}
		year, month, day := now.Date()
		l.buf.Append([]byte(fmt.Sprint(year)))
		l.buf.AppendByte('/')
		l.buf.Append([]byte(fmt.Sprint(int(month))))
		l.buf.AppendByte('/')
		l.buf.Append([]byte(fmt.Sprint(day)))
		l.buf.AppendByte(' ')

		hours, minutes, seconds := now.Clock()
		l.buf.Append([]byte(fmt.Sprint(hours)))
		l.buf.AppendByte(':')
		l.buf.Append([]byte(fmt.Sprint(minutes)))
		l.buf.AppendByte(':')
		l.buf.Append([]byte(fmt.Sprint(seconds)))
		l.buf.AppendByte(' ')

		if l.color {
			l.buf.Append(ColorReset)
		}
	}

	// print data received
	l.buf.Append([]byte(text))

	if len(text) == 0 || text[len(text)-1] != '\n' {
		l.buf.Append([]byte("\n"))
	}

	// flush the output
	l.out.Write(l.buf)
}

func New(out io.Writer) *Logger {
	return &Logger{
		out:       out,
		timestamp: true,
		color:     true,
	}
}
