package ulogger

import (
	"bytes"
	"fmt"
	"strings"
)

// Warning displays a warning message
func (log *Logger) Warning(args ...interface{}) {
	if log.logLevelCode < 1 {
		return
	}
	ch := make(chan int)
	go func(ch chan int) {
		// Create the log that needs to be displayed on stdout
		buf, logStruct := warningPrefix(log)
		log.WarningColor.Fprint(buf, args...)
		log.WarningColor.Print(buf.String())
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

// Warningf displays a warning message
func (log *Logger) Warningf(format string, args ...interface{}) {
	if log.logLevelCode < 1 {
		return
	}
	ch := make(chan int)
	go func(ch chan int) {
		// Create the log that needs to be displayed on stdout
		buf, logStruct := warningPrefix(log)
		log.WarningColor.Fprintf(buf, format, args...)
		// Using Print instead of Printf, since the format string would be taken into account in the Fprintf method
		log.WarningColor.Print(buf.String())

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

// Warningln displays a warning message
func (log *Logger) Warningln(args ...interface{}) {
	if log.logLevelCode < 1 {
		return
	}
	ch := make(chan int)
	go func(ch chan int) {
		// Create the log that needs to be displayed on stdout
		buf, logStruct := warningPrefix(log)
		log.WarningColor.Fprint(buf, args...)
		log.WarningColor.Println(buf.String())
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
func warningPrefix(log *Logger) (*bytes.Buffer, logMessage) {
	buf := new(bytes.Buffer)
	logStruct, timestamp := generateTimestamp()
	logStruct.MessageType = "WARNING"
	log.WarningTimeColor.Fprint(buf, timestamp.Format(timeFormat))
	fmt.Fprint(buf, " ")
	log.WarningMessageTypeColor.Fprint(buf, logStruct.MessageType)
	fmt.Fprint(buf, " ")

	return buf, logStruct
}
