package gcl

var (
	ColorReset  = []byte("\033[0m")
	ColorBlack  = []byte("\033[30m")
	ColorRed    = []byte("\033[31m")
	ColorGreen  = []byte("\033[32m")
	ColorYellow = []byte("\033[33m")
	ColorBlue   = []byte("\033[34m")
	ColorPurple = []byte("\033[35m")
	ColorCyan   = []byte("\033[36m")
	ColorGray   = []byte("\033[37m")
	ColorWhite  = []byte("\033[97m")
)

var (
	ColorBackgroundBlack  = []byte("\u001b[40m")
	ColorBackgroundRed    = []byte(" \u001b[41m")
	ColorBackgroundGreen  = []byte("\u001b[42m")
	ColorBackgroundYellow = []byte("\u001b[43m")
	ColorBackgroundBlue   = []byte(" \u001b[44m")
	ColorBackgroundPurple = []byte("\u001b[45m")
	ColorBackgroundCyan   = []byte("\u001b[46m")
	ColorBackgroundWhite  = []byte("\u001b[47m")
)

var (
	StyleBold      = []byte("\u001b[1m")
	StyleUnderline = []byte("\u001b[4m")
	StyleReversed  = []byte("\u001b[7m")
)

func colorize(data []byte, color []byte) []byte {
	var result []byte
	return append(append(append(result, color...), data...), ColorReset...)
}

func Blue(data []byte) []byte {
	return colorize(data, ColorBlue)
}

func Yellow(data []byte) []byte {
	return colorize(data, ColorYellow)
}

func Red(data []byte) []byte {
	return colorize(data, ColorRed)
}

func Green(data []byte) []byte {
	return colorize(data, ColorGreen)
}

func Cyan(data []byte) []byte {
	return colorize(data, ColorCyan)
}
