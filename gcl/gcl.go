package gcl

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var infoPrefixPlain = []byte("INFO")
var warnPrefixPlain = []byte("WARN")
var errorPrefixPlain = []byte("ERRO")
var fatalPrefixPlain = []byte("FATA")
var successPrefixPlain = []byte("SUCC")

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

type jsonOutput struct {
	Type     string                 `json:"type,omitempty"`
	Message  string                 `json:"message,omitempty"`
	Data     map[string]interface{} `json:"data,omitempty"`
	Time     string                 `json:"time,omitempty"`
	FileInfo string                 `json:"fileInfo,omitempty"`
}

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

func (l *Logger) prepareBuffers(text string) *jsonOutput {
	var (
		timeBuff     Buffer
		fileInfoBuff Buffer
		textBuff     Buffer
		dataBuff     Buffer
	)

	textBuff.Append([]byte(text))

	timeBuff.Append([]byte(time.Now().Format(time.RFC3339)))

	file, line, fn := getFileInfo()
	fileInfoBuff.Append([]byte(fmt.Sprintf("file<%v:%v>@%v", file, line, fn)))

	if jsonBytes, err := json.Marshal(l.fields); err == nil && string(jsonBytes) != "null" {
		dataBuff.Append(jsonBytes)
		l.fields.Reset()
	}

	// Start formatting output

	output := jsonOutput{
		Message:  string(textBuff),
		Time:     string(timeBuff),
		FileInfo: string(fileInfoBuff),
	}

	if len(dataBuff) > 0 {
		json.Unmarshal(dataBuff, &output.Data)
	}

	return &output
}

func (l *Logger) Log(prefix Prefix, text string) {

	l.mu.Lock()
	defer l.mu.Unlock()

	l.buf.Reset()

	output := l.prepareBuffers(text)
	isJson := len(output.Data) > 0
	isColored := l.color && !isJson

	if isColored {
		output.Type = string(prefix.Colored)
	} else {
		output.Type = string(prefix.Plain)
	}

	if isJson {
		// flush data in json format
		if d, err := json.Marshal(output); err == nil {
			l.buf.Append(d[:])
			l.buf.AppendByte('\n')
			l.out.Write(l.buf)
			return
		}
	}

	// create timestamp buff
	if l.timestamp {
		if isColored {
			l.buf.Append(ColorGray)
		}

		t, _ := time.Parse(time.RFC3339, output.Time)
		l.buf.Append(getCurrentTime(t))

		if isColored {
			l.buf.Append(ColorReset)
		}

		l.buf.AppendByte(' ')
	}

	// create prefix buff
	l.buf.Append([]byte(output.Type))
	l.buf.AppendByte(' ')

	// create file info buff
	if l.showFileInfo {
		l.buf.AppendByte('[')
		if isColored {
			l.buf.Append(ColorReset)
			l.buf.Append(StyleUnderline)
		}

		l.buf.Append([]byte(output.FileInfo))

		if isColored {
			l.buf.Append(ColorReset)
		}
		l.buf.AppendByte(']')
		l.buf.AppendByte(' ')
	}

	l.buf.Append([]byte(output.Message))

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
