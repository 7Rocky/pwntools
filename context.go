package pwntools

const (
	DEBUG int = iota
	INFO
	WARNING
	ERROR
	CRITICAL
)

type Context struct {
	LogLevel int
}

var context Context

func init() {
	context.LogLevel = INFO
}

func SetContext(c Context) {
	context = c
}
