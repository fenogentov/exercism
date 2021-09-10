package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Message extracts the message from the provided log line.
func Message(line string) string {
	msg := ""
	if idx := strings.IndexByte(line, ':'); idx >= 0 {
		msg = strings.TrimSpace(line[(idx + 1):])
	}
	return msg
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	msg := Message(line)
	return utf8.RuneCountInString(msg)
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
	lvl := ""
	idx1 := strings.IndexByte(line, '[')
	idx2 := strings.IndexByte(line, ']')
	if idx1 >= 0 && idx2 >= 0 {
		lvl = strings.ToLower(line[(idx1 + 1):idx2])
	}
	return lvl
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	lvl := LogLevel(line)
	msg := Message(line)
	return fmt.Sprintf("%s (%s)", msg, lvl)
}
