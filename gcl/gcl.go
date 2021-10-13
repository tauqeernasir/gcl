package gcl

import (
	"io"
	"sync"
)

var infoPrefixPlain = []byte("[INFO] ")

type Prefix struct {
	Plain   []byte
	Colored []byte
}

var (
	InfoPrefix = Prefix{
		Plain:   infoPrefixPlain,
		Colored: Blue(infoPrefixPlain),
	}
)

type Logger struct {
	mu        sync.RWMutex
	timestamp bool
	buf       Buffer
	out       io.Writer
}

func (l *Logger) Info(text string) {
	// fmt.Printf("%v%v%v%v\n", InfoPrefix, colorBlue, text, colorReset)
	l.Log(InfoPrefix, text)
}

func (l *Logger) WithTimestamp(t bool) *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.timestamp = t
	return l
}

func (l *Logger) Log(prefix Prefix, text string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.buf.Reset()

	l.buf.Append(prefix.Colored)

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
		out: out,
	}
}
