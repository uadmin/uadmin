// +build !windows

package colors

const (
	// FGBlackB Forground Black Bold
	FGBlackB = "\x1b[30;1m"
	// FGRedB Forground Red Bold
	FGRedB = "\x1b[31;1m"
	// FGGreenB Forground Green Bold
	FGGreenB = "\x1b[32;1m"
	// FGYellowB Forground Yellow Bold
	FGYellowB = "\x1b[33;1m"
	// FGBlueB Forground Blue Bold
	FGBlueB = "\x1b[34;1m"
	// FGMagentaB Forground Magenta Bold
	FGMagentaB = "\x1b[35;1m"
	// FGCyanB Forground Cyan Bold
	FGCyanB = "\x1b[36;1m"
	// FGWhiteB Forground White Bold
	FGWhiteB = "\x1b[37;1m"
	// FGBlack Forground Black
	FGBlack = "\x1b[30m"
	// FGRed Forground Red
	FGRed = "\x1b[31m"
	// FGGreen Forground Green
	FGGreen = "\x1b[32m"
	// FGYellow Forground Yellow
	FGYellow = "\x1b[33m"
	// FGBlue Forground Blue
	FGBlue = "\x1b[34m"
	// FGMagenta Forground Magenta
	FGMagenta = "\x1b[35m"
	// FGCyan Forground Cyan
	FGCyan = "\x1b[36m"
	// FGWhite Forground White
	FGWhite = "\x1b[37m"
	// FGNormal Forground Reset to Normal
	FGNormal = "\x1b[0m"
	// BGBlack Background Black
	BGBlack = "\x1b[40m"
	// BGRed Background Red
	BGRed = "\x1b[41m"
	// BGGreen Background Green
	BGGreen = "\x1b[42m"
	// BGYellow Background Yellow
	BGYellow = "\x1b[43m"
	// BGBlue Background Blue
	BGBlue = "\x1b[44m"
	// BGMagenta Background Magenta
	BGMagenta = "\x1b[45m"
	// BGCyan Background Cyan
	BGCyan = "\x1b[46m"
	// BGWhite Background White
	BGWhite = "\x1b[47m"
)

// OK CLI Display
const OK = "[   " + FGGreenB + "OK" + FGNormal + "   ]   "

// Working CLI Display
const Working = "[ " + FGMagentaB + "WORKING" + FGNormal + "]   "

// Warning CLI Display
const Warning = "[ " + FGYellowB + "WARNING" + FGNormal + "]   "

// Error CLI Display
const Error = "[  " + FGRedB + "ERROR" + FGNormal + " ]   "

// Debug CLI Display
const Debug = "[  " + FGCyanB + "DEBUG" + FGNormal + " ]   "

// Info CLI Display
const Info = "[  " + FGBlueB + "INFO" + FGNormal + "  ]   "

// Critical CLI Display
const Critical = "[" + FGRedB + "CRITICAL" + FGNormal + "]   "

// Alert CLI Display
const Alert = "[  " + FGRedB + "ALERT" + FGNormal + " ]   "

// Emergency CLI Display
const Emergency = "[  " + FGRedB + "EMERG" + FGNormal + " ]   "
