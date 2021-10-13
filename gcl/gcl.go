package gcl

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var infoPrefixPlain = []byte("INFO \u2591 ")
var warnPrefixPlain = []byte("WARN \u2591 ")
var errorPrefixPlain = []byte("ERRO \u2591 ")
var fatalPrefixPlain = []byte("FATA \u2591 ")
var successPrefixPlain = []byte("SUCC \u2591 ")

type Prefix struct {
	Plain   []byte
	Colored []byte
}

var (
	InfoPrefix = Prefix{
		Plain:   infoPrefixPlain,
		Colored: Cyan(infoPrefixPlain),
	}

	WarnPrefix = Prefix{
		Plain:   warnPrefixPlain,
		Colored: Yellow(warnPrefixPlain),
	}

	ErrorPrefix = Prefix{
		Plain:   errorPrefixPlain,
		Colored: Red(errorPrefixPlain),
	}

	FatalPrefix = Prefix{
		Plain:   fatalPrefixPlain,
		Colored: Red(fatalPrefixPlain),
	}

	SuccessPrefix = Prefix{
		Plain:   successPrefixPlain,
		Colored: Green(successPrefixPlain),
	}
)

type Logger struct {
	mu           sync.RWMutex
	color        bool
	timestamp    bool
	buf          Buffer
	out          io.Writer
	showFileInfo bool
	// fields to display with next out call
	fields Fields
}

type Fields map[string]interface{}

func (f *Fields) Reset() {
	*f = nil
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

func (l *Logger) Fatal(text string) {
	l.Log(FatalPrefix, text)
	os.Exit(1)
}

func (l *Logger) Fatalf(text string, args ...interface{}) {
	l.Log(FatalPrefix, fmt.Sprintf(text, args...))
	os.Exit(1)
}

func (l *Logger) Success(text string) {
	l.Log(SuccessPrefix, text)
}

func (l *Logger) WithFields(fields Fields) *Logger {
	l.fields = fields
	return l
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

func (l *Logger) WithFileInfo() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.showFileInfo = true
	return l
}

func (l *Logger) WithoutFileInfo() *Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.showFileInfo = false
	return l
}

func (l *Logger) Log(prefix Prefix, text string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()

	l.buf.Reset()

	if l.timestamp {
		if l.color {
			l.buf.Append(ColorGray)
		}

		formattedTimeBytes := getCurrentTime(now)

		l.buf.Append(formattedTimeBytes)

		if l.color {
			l.buf.Append(ColorReset)
		}

		l.buf.AppendByte(' ')
	}

	if l.color {
		l.buf.Append(prefix.Colored)
	} else {
		l.buf.Append(prefix.Plain)
	}

	if l.showFileInfo {
		l.buf.AppendByte('[')
		if l.color {
			l.buf.Append(ColorReset)
			l.buf.Append(StyleUnderline)
		}

		file, line, fn := getFileInfo()
		l.buf.Append([]byte(fmt.Sprintf("file<%v:%v>@%v", file, line, fn)))

		if l.color {
			l.buf.Append(ColorReset)
		}
		l.buf.AppendByte(']')
		l.buf.AppendByte(' ')
	}

	// print data received
	l.buf.Append([]byte(text))

	if jsonBytes, err := json.Marshal(l.fields); err == nil && string(jsonBytes) != "null" {
		l.buf.AppendByte(' ')
		l.buf.Append([]byte(jsonBytes))
		l.fields.Reset()
	}

	if len(text) == 0 || text[len(text)-1] != '\n' {
		l.buf.Append([]byte("\n"))
	}

	// flush the output
	l.out.Write(l.buf)
}

func NewLogger(out io.Writer) *Logger {
	return &Logger{
		out:       out,
		timestamp: true,
		color:     true,
	}
}
