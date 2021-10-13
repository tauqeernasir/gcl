package gcl

var (
	colorReset  = []byte("\033[0m")
	colorBlack  = []byte("\033[30m")
	colorRed    = []byte("\033[31m")
	colorGreen  = []byte("\033[32m")
	colorYellow = []byte("\033[33m")
	colorBlue   = []byte("\033[34m")
	colorPurple = []byte("\033[35m")
	colorCyan   = []byte("\033[36m")
	colorGray   = []byte("\033[37m")
	colorWhite  = []byte("\033[97m")
)

var (
	colorBackgroundBlack   = []byte("\u001b[40m")
	colorBackgroundRed     = []byte(" \u001b[41m")
	colorBackgroundGreen   = []byte("\u001b[42m")
	colorBackgroundYellow  = []byte("\u001b[43m")
	colorBackgroundBlue    = []byte(" \u001b[44m")
	colorBackgroundMagenta = []byte("\u001b[45m")
	colorBackgroundCyan    = []byte("\u001b[46m")
	colorBackgroundWhite   = []byte("\u001b[47m")
)

var (
	styleBold      = []byte("\u001b[1m")
	styleUnderline = []byte("\u001b[4m")
	styleReversed  = []byte("\u001b[7m")
)

func colorize(data []byte, color []byte) []byte {
	var result []byte
	return append(append(append(result, color...), data...), colorReset...)
}

func Blue(data []byte) []byte {
	return colorize(data, colorBlue)
}
