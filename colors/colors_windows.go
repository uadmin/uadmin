// +build windows

package colors

const (
	// FGBlackB Forground Black Bold
	FGBlackB = ""
	// FGRedB Forground Red Bold
	FGRedB = ""
	// FGGreenB Forground Green Bold
	FGGreenB = ""
	// FGYellowB Forground Yellow Bold
	FGYellowB = ""
	// FGBlueB Forground Blue Bold
	FGBlueB = ""
	// FGMagentaB Forground Magenta Bold
	FGMagentaB = ""
	// FGCyanB Forground Cyan Bold
	FGCyanB = ""
	// FGWhiteB Forground White Bold
	FGWhiteB = ""
	// FGBlack Forground Black
	FGBlack = ""
	// FGRed Forground Red
	FGRed = ""
	// FGGreen Forground Green
	FGGreen = ""
	// FGYellow Forground Yellow
	FGYellow = ""
	// FGBlue Forground Blue
	FGBlue = ""
	// FGMagenta Forground Magenta
	FGMagenta = ""
	// FGCyan Forground Cyan
	FGCyan = ""
	// FGWhite Forground White
	FGWhite = ""
	// FGNormal Forground Reset to Normal
	FGNormal = ""
	// BGBlack Background Black
	BGBlack = ""
	// BGRed Background Red
	BGRed = ""
	// BGGreen Background Green
	BGGreen = ""
	// BGYellow Background Yellow
	BGYellow = ""
	// BGBlue Background Blue
	BGBlue = ""
	// BGMagenta Background Magenta
	BGMagenta = ""
	// BGCyan Background Cyan
	BGCyan = ""
	// BGWhite Background White
	BGWhite = ""
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
