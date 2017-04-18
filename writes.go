package ulogger

import (
	"bytes"
	"fmt"
	"time"
)

// writeToPrint writes the logging message to a string and returns it
func (log *Logger) writeToPrint(buf *bytes.Buffer, args ...interface{}) string {
	fmt.Fprint(buf, time.Now().Format(timeFormat), "\t", args)
	return buf.String()
}

// writeToPrintf writes the logging message to a string and returns it
func (log *Logger) writeToPrintf(buf *bytes.Buffer, format string, args ...interface{}) string {
	fmt.Fprintf(buf, format, args...)
	return buf.String()
}

// writeToPrintln writes the logging message to a string and returns it
func (log *Logger) writeToPrintln(buf *bytes.Buffer, args ...interface{}) string {
	fmt.Fprintln(buf, args...)
	return buf.String()
}
