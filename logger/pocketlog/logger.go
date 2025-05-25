package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// LOgger is used to log inmsgion
type Logger struct {
	threshold            Level
	output               io.Writer
	msgFormattingOptions []LogMessageOption
}

// New returns you a logger, ready to log at the required threshold.
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{threshold: threshold, output: os.Stdout}

	for _, configFunc := range opts {
		if configFunc != nil {
			configFunc(lgr)
		}

	}
	return lgr
}

func (l *Logger) logf(msg string, args ...any) {

	if l.output == nil {
		l.output = os.Stdout
	}

	_, _ = fmt.Fprintf(l.output, msg+"\n", args...)
}

func (l *Logger) Logf(lvl Level, msg string, args ...any) {
	if l.threshold > lvl {
		return
	}

	if len(l.msgFormattingOptions) > 0 {
		for _, logMsgConfigFunc := range l.msgFormattingOptions {
			msg = logMsgConfigFunc(msg, lvl)
		}
	}

	l.logf(msg, args...)
}

// Debugf msgs and prints a message (like fmt.Printf) if the log level is debug or higher.
func (l *Logger) Debugf(msg string, args ...any) {
	l.Logf(LevelDebug, msg, args...)

}

// Infof msgs and prints a message (like fmt.Printf) if the log level is info or higher.
func (l *Logger) Infof(msg string, args ...any) {
	l.Logf(LevelInfo, msg, args...)
}

// Errorf msgs and prints a message (like fmt.Printf) if the log level is error or higher.
func (l *Logger) Errorf(msg string, args ...any) {
	l.Logf(LevelError, msg, args...)
}
