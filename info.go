package ulogger

import (
	"bytes"
	"fmt"
	"strings"
)

// Info displays a non-fatal log message
func (log *Logger) Info(args ...interface{}) {
	if log.logLevelCode < 1 {
		return
	}
	ch := make(chan int)
	go func(ch chan int) {
		// Create the log that needs to be displayed on stdout
		buf, logStruct := infoPrefix(log)
		log.InfoColor.Fprint(buf, args...)
		log.InfoColor.Print(buf.String())
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

// Infof displays a non-fatal log message according to the format string
func (log *Logger) Infof(format string, args ...interface{}) {
	if log.logLevelCode < 1 {
		return
	}
	ch := make(chan int)
	go func(ch chan int) {
		// Create the log that needs to be displayed on stdout
		buf, logStruct := infoPrefix(log)
		log.InfoColor.Fprintf(buf, format, args...)
		// Using Print instead of Printf, since the format string would be taken into account in the Fprintf method
		log.InfoColor.Print(buf.String())

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

// Infoln displays a non-fatal log message
func (log *Logger) Infoln(args ...interface{}) {
	if log.logLevelCode < 1 {
		return
	}
	ch := make(chan int)
	go func(ch chan int) {
		// Create the log that needs to be displayed on stdout
		buf, logStruct := infoPrefix(log)
		log.InfoColor.Fprint(buf, args...)
		log.InfoColor.Println(buf.String())
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
func infoPrefix(log *Logger) (*bytes.Buffer, logMessage) {
	buf := new(bytes.Buffer)
	logStruct, timestamp := generateTimestamp()
	logStruct.MessageType = "INFO"
	log.InfoTimeColor.Fprint(buf, timestamp.Format(timeFormat))
	fmt.Fprint(buf, " ")
	log.InfoMessageTypeColor.Fprint(buf, logStruct.MessageType)
	fmt.Fprint(buf, " ")

	return buf, logStruct
}
