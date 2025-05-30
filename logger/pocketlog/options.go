package pocketlog

import (
	"fmt"
	"io"
	"time"
)

type Option func(*Logger)

type LogMessageOption func(string, Level) string

// WithOutput returns a configuration function that sets the output of logs
func WithOutput(output io.Writer) Option {
	return func(lgr *Logger) {
		lgr.output = output
	}
}

// AddLogMessageOptions adds log message configuring options to logger object
func AddLogMessageOptions(funcs ...LogMessageOption) Option {
	return func(lgr *Logger) {
		lgr.msgFormattingOptions = funcs
	}
}

// Adds log level based prefix to each log, like [ERROR], [INFO] etc
func AddPrefixBasedOnLogLevel() LogMessageOption {
	return func(msg string, lvl Level) string {
		if logPrefix, ok := logPrefixMap[lvl]; ok {
			return fmt.Sprintf("[%s] %s", logPrefix, msg)
		}
		return msg
	}
}

// Adds date to each log 
func AddDate() LogMessageOption {
	return func(msg string, lvl Level) string {
		return fmt.Sprintf("%s| %s", time.Now().Format(time.RFC3339), msg)
	}
}
