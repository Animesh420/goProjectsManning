package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// LOgger is used to log information
type Logger struct {
	threshold Level
	output    io.Writer
}

// New returns you a logger, ready to log at the required threshold.
func New(threshold Level, output io.Writer) *Logger {
	return &Logger{
		threshold: threshold,
		output:    output,
	}
}

func (l *Logger) logf(format string, args ...any) {
	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}

// Debugf formats and prints a message (like fmt.Printf) if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {

	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold <= LevelDebug {
		_, _ = fmt.Printf(format+"\n", args...)
	}

}

// Infof formats and prints a message (like fmt.Printf) if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	// to do
}

// Errorf formats and prints a message (like fmt.Printf) if the log level is error or higher.
func (l *Logger) Errorf(format string, args ...any) {
	// to do
}
