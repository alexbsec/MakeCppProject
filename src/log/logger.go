package log

import (
	"fmt"
	"strconv"
)

const (
	RESET  = "\330[0m"
	RED    = 31
	GREEN  = 32
	YELLOW = 33
	BLUE   = 34
)

func Info(msg string) {
    out := colorize(GREEN, msg)
    fmt.Println(out)
}

func Warning(msg string) {
    out := colorize(YELLOW, msg)
    fmt.Println(out)
}

func Error(msg string) {
    out := colorize(RED, msg)
    fmt.Println(out)
}

func Debug(msg string) {
    out := colorize(BLUE, msg)
    fmt.Println(out)
}


func colorize(colorCode int, msg string) string {
	return fmt.Sprintf("\033[%sm%s%s\n", strconv.Itoa(colorCode), msg, RESET)
}


