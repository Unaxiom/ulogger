package ulogger

import (
	"bytes"
	"fmt"
)

// Debug displays a debugging message useful in development environment
func (log *Logger) Debug(args ...interface{}) {
	if log.logLevelCode < 1 {
		return
	}
	ch := make(chan int)
	go func(ch chan int) {
		write(debugPrefix, log, log.DebugColor, ch, args...)
	}(ch)
	<-ch
}

// Debugf displays a debugging message
func (log *Logger) Debugf(format string, args ...interface{}) {
	if log.logLevelCode < 1 {
		return
	}
	ch := make(chan int)
	go func(ch chan int) {
		writef(debugPrefix, log, log.DebugColor, ch, format, args...)
	}(ch)
	<-ch
}

// Debugln displays a debugging message
func (log *Logger) Debugln(args ...interface{}) {
	if log.logLevelCode < 1 {
		return
	}
	ch := make(chan int)
	go func(ch chan int) {
		writeln(debugPrefix, log, log.DebugColor, ch, args...)
	}(ch)
	<-ch
}

// Returns a string, along with a logMessage after prefixing the timestamp and the type of log
func debugPrefix(log *Logger) (*bytes.Buffer, logMessage) {
	buf := new(bytes.Buffer)
	logStruct, timestamp := generateTimestamp()
	logStruct.MessageType = "DEBUG"
	log.DebugTimeColor.Fprint(buf, timestamp.Format(timeFormat))
	fmt.Fprint(buf, " ")
	log.DebugMessageTypeColor.Fprint(buf, logStruct.MessageType)
	fmt.Fprint(buf, " ")
	return buf, logStruct
}
