package ulogger

import (
	"bytes"
	"fmt"
)

// Warning displays a warning message
func (log *Logger) Warning(args ...interface{}) {
	if log.logLevelCode > 3 {
		if log.RemoteAvailable {
			// Create the logMessage struct here
			logStruct, _ := generateTimestamp("WARNING")
			ch := make(chan int)
			go sendLogMessageFromWrite(logStruct, ch, args...)
			<-ch
		}
		return
	}
	ch := make(chan int)
	go func(ch chan int) {
		write(warningPrefix, log, log.WarningColor, ch, args...)
	}(ch)
	<-ch
}

// Warningf displays a warning message
func (log *Logger) Warningf(format string, args ...interface{}) {
	if log.logLevelCode > 3 {
		if log.RemoteAvailable {
			// Create the logMessage struct here
			logStruct, _ := generateTimestamp("WARNING")
			ch := make(chan int)
			go sendLogMessageFromWritef(logStruct, ch, format, args...)
			<-ch
		}
		return
	}
	ch := make(chan int)
	go func(ch chan int) {
		writef(warningPrefix, log, log.WarningColor, ch, format, args...)
	}(ch)
	<-ch
}

// Warningln displays a warning message
func (log *Logger) Warningln(args ...interface{}) {
	if log.logLevelCode > 3 {
		if log.RemoteAvailable {
			// Create the logMessage struct here
			logStruct, _ := generateTimestamp("WARNING")
			ch := make(chan int)
			go sendLogMessageFromWriteln(logStruct, ch, args...)
			<-ch
		}
		return
	}
	ch := make(chan int)
	go func(ch chan int) {
		writeln(warningPrefix, log, log.WarningColor, ch, args...)
	}(ch)
	<-ch
}

// Returns a string, along with a logMessage after prefixing the timestamp and the type of log
func warningPrefix(log *Logger) (*bytes.Buffer, logMessage) {
	buf := new(bytes.Buffer)
	logStruct, timestamp := generateTimestamp("WARNING")
	logStruct.OrganizationName = log.OrganizationName
	logStruct.ApplicationName = log.ApplicationName
	log.WarningTimeColor.Fprint(buf, timestamp.Format(timeFormat))
	fmt.Fprint(buf, " ")
	log.WarningMessageTypeColor.Fprint(buf, logStruct.MessageType)
	fmt.Fprint(buf, " ")
	return buf, logStruct
}
