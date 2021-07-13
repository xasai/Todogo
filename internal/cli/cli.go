package cli

import (
	"os"
)

var (
	RED = ""
	GREEN = ""
	YELL = ""
	PURP = ""
	PINK = ""
	CYAN = ""
	RES = ""
)


func init() {
	if os.Getenv("COLORTERM") == "truecolor" {
		RED = "\033[1;31m"
		GREEN = "\033[1;32m"
		YELL = "\033[1;33m"
		PURP = "\033[1;34m"
		PINK = "\033[1;35m"
		CYAN = "\033[1;36m"
		RES = "\033[1;0m"
	}
}
