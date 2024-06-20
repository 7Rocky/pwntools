package pwntools

import (
	"fmt"
	"log"
	"strings"
)

const (
	RESET = "\x1b[0m"

	BLUE   = "\x1b[0;34;1m"
	GREEN  = "\x1b[0;32;1m"
	RED    = "\x1b[0;31;1m"
	YELLOW = "\x1b[0;33;1m"

	RED_BG = "\x1b[0;41m"
)

type ProgressBar struct {
	Title      string
	LastStatus string
}

func init() {
	log.Default().SetFlags(0)
}

func padding(next, prev string) string {
	if len(next) < len(prev) {
		return strings.Repeat(" ", len(prev)-len(next))
	}

	return ""
}

func Progress(title string) ProgressBar {
	if context.LogLevel <= INFO {
		fmt.Printf("[%s*%s] %s\r", BLUE, RESET, title)
		return ProgressBar{Title: title, LastStatus: ""}
	}

	return ProgressBar{}
}

func (p *ProgressBar) Status(status string) {
	if context.LogLevel <= INFO {
		fmt.Printf("[%s*%s] %s: %s%s\r", BLUE, RESET, p.Title, status, padding(status, p.LastStatus))
		p.LastStatus = status
	}
}

func (p *ProgressBar) Success(status string) {
	if context.LogLevel <= INFO {
		fmt.Printf("[%s+%s] %s: %s%s\n", GREEN, RESET, p.Title, status, padding(status, p.LastStatus))
	}
}

func (p *ProgressBar) Failure(status string) {
	if context.LogLevel <= INFO {
		fmt.Printf("[%s-%s] %s: %s%s\n", RED, RESET, p.Title, status, padding(status, p.LastStatus))
	}
}

func Critical(format string, v ...any) {
	if context.LogLevel <= CRITICAL {
		log.Printf("[%sCRITICAL%s] "+format, append([]any{RED_BG, RESET}, v...)...)
	}
}

func Debug(format string, v ...any) {
	if context.LogLevel <= DEBUG {
		log.Printf("[%sDEBUG%s] "+format, append([]any{RED, RESET}, v...)...)
		panic(fmt.Sprintf(format, v...))
	}
}

func Error(format string, v ...any) {
	if context.LogLevel <= ERROR {
		log.Printf("[%sERROR%s] "+format, append([]any{RED_BG, RESET}, v...)...)
		panic(fmt.Sprintf(format, v...))
	}
}

func Failure(format string, v ...any) {
	if context.LogLevel <= ERROR {
		log.Printf("[%s-%s] "+format, append([]any{RED, RESET}, v...)...)
	}
}

func Info(format string, v ...any) {
	if context.LogLevel <= INFO {
		log.Printf("[%s*%s] "+format, append([]any{BLUE, RESET}, v...)...)
	}
}

func Success(format string, v ...any) {
	if context.LogLevel <= INFO {
		log.Printf("[%s+%s] "+format, append([]any{GREEN, RESET}, v...)...)
	}
}

func Warning(format string, v ...any) {
	if context.LogLevel <= WARNING {
		log.Printf("[%s!%s] "+format, append([]any{YELLOW, RESET}, v...)...)
	}
}
