package ulogger

import (
	"bytes"
	"fmt"
	"strings"
)

// Debug displays a debugging message useful in development environment
func (log *Logger) Debug(args ...interface{}) {
	if log.logLevelCode < 1 {
		return
	}
	ch := make(chan int)
	go func(ch chan int) {
		// Create the log that needs to be displayed on stdout
		buf, logStruct := debugPrefix(log)
		log.DebugColor.Fprint(buf, args...)
		log.DebugColor.Print(buf.String())
		go func() {
			// Create the log message that needs to be sent to the server, only if the RemoteAvailable flag is set
			if !log.RemoteAvailable {
				ch <- 1
				return
			}
			buf.Reset()
			fmt.Fprint(buf, args...)
			logStruct.MessageContent = strings.TrimSpace(buf.String())
			go func() {
				// Send the actual message here
				sendLogMessage(logStruct)
				ch <- 1
			}()
		}()
		ch <- 1
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
		// Create the log that needs to be displayed on stdout
		buf, logStruct := debugPrefix(log)
		log.DebugColor.Fprintf(buf, format, args...)
		// Using Print instead of Printf, since the format string would be taken into account in the Fprintf method
		log.DebugColor.Print(buf.String())

		go func() {
			// Create the log message that needs to be sent to the server, only if the RemoteAvailable flag is set
			if !log.RemoteAvailable {
				ch <- 1
				return
			}
			buf.Reset()
			fmt.Fprintf(buf, format, args...)
			logStruct.MessageContent = strings.TrimSpace(buf.String())
			go func() {
				sendLogMessage(logStruct)
				ch <- 1
			}()
		}()
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
		// Create the log that needs to be displayed on stdout
		buf, logStruct := debugPrefix(log)
		log.DebugColor.Fprint(buf, args...)
		log.DebugColor.Println(buf.String())
		go func() {
			// Create the log message that needs to be sent to the server, only if the RemoteAvailable flag is set
			if !log.RemoteAvailable {
				ch <- 1
				return
			}
			buf.Reset()
			fmt.Fprintln(buf, args...)
			logStruct.MessageContent = strings.TrimSpace(buf.String())
			go func() {
				// Send the actual message here
				sendLogMessage(logStruct)
				ch <- 1
			}()
		}()
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
