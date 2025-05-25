package pocketlog

// Level represents an available logging level
type Level byte

const (
	// LevelDebug represents the lowest level of log, used for debugging purposes
	LevelDebug Level = iota
	// LevelInfo represents a logging level that contains information deemed valuable
	LevelInfo
	// LevelError represents the highest logging level, only to be used to trace errors
	LevelError
)

var logPrefixMap = map[Level]string{
	LevelDebug: "DEBUG",
	LevelInfo:  "INFO",
	LevelError: "ERROR",
}
