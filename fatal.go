package ulogger

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// Fatal displays a message and crashes the program
func (log *Logger) Fatal(args ...interface{}) {
	if log.logLevelCode < 1 {
		return
	}
	ch := make(chan int)
	go func(ch chan int) {
		// Create the log that needs to be displayed on stdout
		buf, logStruct := fatalPrefix(log)
		log.FatalColor.Fprint(buf, args...)
		log.FatalColor.Print(buf.String())
		go func() {
			// Create the log message that needs to be sent to the server, only if the RemoteAvailable flag is set
			if !log.RemoteAvailable {
				ch <- 1
				os.Exit(1)
				return
			}
			buf.Reset()
			fmt.Fprint(buf, args...)
			logStruct.MessageContent = strings.TrimSpace(buf.String())
			go func() {
				// Send the actual message here
				sendLogMessage(logStruct)
				os.Exit(1)
				ch <- 1
			}()
		}()
		ch <- 1
	}(ch)
	<-ch
}

// Fatalf displays a message and crashes the program
func (log *Logger) Fatalf(format string, args ...interface{}) {
	if log.logLevelCode < 1 {
		return
	}
	ch := make(chan int)
	go func(ch chan int) {
		// Create the log that needs to be displayed on stdout
		buf, logStruct := fatalPrefix(log)
		log.FatalColor.Fprintf(buf, format, args...)
		// Using Print instead of Printf, since the format string would be taken into account in the Fprintf method
		log.FatalColor.Print(buf.String())

		go func() {
			// Create the log message that needs to be sent to the server, only if the RemoteAvailable flag is set
			if !log.RemoteAvailable {
				ch <- 1
				os.Exit(1)
			}
			buf.Reset()
			fmt.Fprintf(buf, format, args...)
			logStruct.MessageContent = strings.TrimSpace(buf.String())
			go func() {
				sendLogMessage(logStruct)
				ch <- 1
				os.Exit(1)
			}()
		}()
	}(ch)
	<-ch
}

// Fatalln displays a message and crashes the program
func (log *Logger) Fatalln(args ...interface{}) {
	if log.logLevelCode < 1 {
		return
	}
	ch := make(chan int)
	go func(ch chan int) {
		// Create the log that needs to be displayed on stdout
		buf, logStruct := fatalPrefix(log)
		log.FatalColor.Fprint(buf, args...)
		log.FatalColor.Println(buf.String())
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
func fatalPrefix(log *Logger) (*bytes.Buffer, logMessage) {
	buf := new(bytes.Buffer)
	logStruct, timestamp := generateTimestamp()
	logStruct.MessageType = "FATAL"
	log.FatalTimeColor.Fprint(buf, timestamp.Format(timeFormat))
	fmt.Fprint(buf, " ")
	log.FatalMessageTypeColor.Fprint(buf, logStruct.MessageType)
	fmt.Fprint(buf, " ")

	return buf, logStruct
}
