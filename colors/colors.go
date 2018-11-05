// +build !windows

package colors

const FG_BLACK_B = "\x1b[30;1m"
const FG_RED_B = "\x1b[31;1m"
const FG_GREEN_B = "\x1b[32;1m"
const FG_YELLOW_B = "\x1b[33;1m"
const FG_BLUE_B = "\x1b[34;1m"
const FG_MAGENTA_B = "\x1b[35;1m"
const FG_CYAN_B = "\x1b[36;1m"
const FG_WHITE_B = "\x1b[37;1m"

const FG_BLACK = "\x1b[30m"
const FG_RED = "\x1b[31m"
const FG_GREEN = "\x1b[32m"
const FG_YELLOW = "\x1b[33m"
const FG_BLUE = "\x1b[34m"
const FG_MAGENTA = "\x1b[35m"
const FG_CYAN = "\x1b[36m"
const FG_WHITE = "\x1b[37m"
const FG_NORMAL = "\x1b[0m"

const BG_BLACK = "\x1b[40m"
const BG_RED = "\x1b[41m"
const BG_GREEN = "\x1b[42m"
const BG_YELLOW = "\x1b[43m"
const BG_BLUE = "\x1b[44m"
const BG_MAGENTA = "\x1b[45m"
const BG_CYAN = "\x1b[46m"
const BG_WHITE = "\x1b[47m"

const OK = "[   " + FG_GREEN_B + "OK" + FG_NORMAL + "   ]   "
const Working = "[ " + FG_MAGENTA_B + "WORKING" + FG_NORMAL + "]   "
const Warning = "[ " + FG_YELLOW_B + "WARNING" + FG_NORMAL + "]   "
const Error = "[  " + FG_RED_B + "ERROR" + FG_NORMAL + " ]   "
const Debug = "[  " + FG_CYAN_B + "DEBUG" + FG_NORMAL + " ]   "
const Info = "[  " + FG_BLUE_B + "INFO" + FG_NORMAL + "  ]   "
