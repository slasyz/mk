package logger

import (
	"fmt"
	"io"
)

type Logger struct {
	inner io.Writer
}

func New(inner io.Writer) *Logger {
	return &Logger{inner}
}

func (l *Logger) Command(cmd string, args []string) {
	if l.inner == nil {
		return
	}

	for i, arg := range args {
		fmt.Fprintf(l.inner, "-> $%d = %s\n", i+1, arg)
	}
	fmt.Fprintln(l.inner, "=>", cmd)
}
