package pwntools

import (
	"fmt"
	"log"
	"strings"
)

const (
	reset = "\x1b[0m"

	blue   = "\x1b[0;34;1m"
	green  = "\x1b[0;32;1m"
	red    = "\x1b[0;31;1m"
	yellow = "\x1b[0;33;1m"

	redBg = "\x1b[0;41m"
)

type progressBar struct {
	title      string
	lastStatus string
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

func Progress(title string) progressBar {
	if context.LogLevel <= INFO {
		fmt.Printf("[%s*%s] %s\r", blue, reset, title)
		return progressBar{title: title, lastStatus: ""}
	}

	return progressBar{}
}

func (p *progressBar) Status(status string) {
	if context.LogLevel <= INFO {
		fmt.Printf("[%s*%s] %s: %s%s\r", blue, reset, p.title, status, padding(status, p.lastStatus))
		p.lastStatus = status
	}
}

func (p *progressBar) Success(status string) {
	if context.LogLevel <= INFO {
		fmt.Printf("[%s+%s] %s: %s%s\n", green, reset, p.title, status, padding(status, p.lastStatus))
	}
}

func (p *progressBar) Failure(status string) {
	if context.LogLevel <= INFO {
		fmt.Printf("[%s-%s] %s: %s%s\n", red, reset, p.title, status, padding(status, p.lastStatus))
	}
}

func Critical(format string, v ...any) {
	if context.LogLevel <= CRITICAL {
		log.Printf("[%sCRITICAL%s] "+format, append([]any{redBg, reset}, v...)...)
	}
}

func Debug(format string, v ...any) {
	if context.LogLevel <= DEBUG {
		log.Printf("[%sDEBUG%s] "+format, append([]any{red, reset}, v...)...)
		panic(fmt.Sprintf(format, v...))
	}
}

func Error(format string, v ...any) {
	if context.LogLevel <= ERROR {
		log.Fatalf("[%sERROR%s] "+format, append([]any{redBg, reset}, v...)...)
	}
}

func Failure(format string, v ...any) {
	if context.LogLevel <= ERROR {
		log.Printf("[%s-%s] "+format, append([]any{red, reset}, v...)...)
	}
}

func Info(format string, v ...any) {
	if context.LogLevel <= INFO {
		log.Printf("[%s*%s] "+format, append([]any{blue, reset}, v...)...)
	}
}

func Success(format string, v ...any) {
	if context.LogLevel <= INFO {
		log.Printf("[%s+%s] "+format, append([]any{green, reset}, v...)...)
	}
}

func Warning(format string, v ...any) {
	if context.LogLevel <= WARNING {
		log.Printf("[%s!%s] "+format, append([]any{yellow, reset}, v...)...)
	}
}
