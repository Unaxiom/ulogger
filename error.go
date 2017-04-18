package ulogger

import (
	"bytes"
	"fmt"
)

// Error displays an error message
func (log *Logger) Error(args ...interface{}) {
	if log.logLevelCode < 1 {
		return
	}
	ch := make(chan int)
	go func(ch chan int) {
		write(errorPrefix, log, log.ErrorColor, ch, args...)
	}(ch)
	<-ch
}

// Errorf displays an error message
func (log *Logger) Errorf(format string, args ...interface{}) {
	if log.logLevelCode < 1 {
		return
	}
	ch := make(chan int)
	go func(ch chan int) {
		writef(errorPrefix, log, log.ErrorColor, ch, format, args...)
	}(ch)
	<-ch
}

// Errorln displays an error message
func (log *Logger) Errorln(args ...interface{}) {
	if log.logLevelCode < 1 {
		return
	}
	ch := make(chan int)
	go func(ch chan int) {
		writeln(errorPrefix, log, log.ErrorColor, ch, args...)
	}(ch)
	<-ch
}

// Returns a string, along with a logMessage after prefixing the timestamp and the type of log
func errorPrefix(log *Logger) (*bytes.Buffer, logMessage) {
	buf := new(bytes.Buffer)
	logStruct, timestamp := generateTimestamp()
	logStruct.MessageType = "ERROR"
	log.ErrorTimeColor.Fprint(buf, timestamp.Format(timeFormat))
	fmt.Fprint(buf, " ")
	log.ErrorMessageTypeColor.Fprint(buf, logStruct.MessageType)
	fmt.Fprint(buf, " ")
	return buf, logStruct
}
