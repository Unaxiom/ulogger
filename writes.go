package ulogger

import (
	"bytes"
	"fmt"
	"strings"

	"os"

	"github.com/fatih/color"
)

type prefixerSignature func(log *Logger) (*bytes.Buffer, logMessage)

// write is applicable for all simple logs
func write(prefixFunc prefixerSignature, log *Logger, clr *color.Color, ch chan int, args ...interface{}) {
	// Create the log that needs to be displayed on stdout
	buf, logStruct := prefixFunc(log)
	clr.Fprint(buf, args...)
	clr.Print(buf.String())
	go func() {
		// Create the log message that needs to be sent to the server, only if the RemoteAvailable flag is set
		if !log.RemoteAvailable {
			ch <- 1
			if logStruct.MessageType == "FATAL" {
				os.Exit(1)
			}
			return
		}
		buf.Reset()
		fmt.Fprint(buf, args...)
		logStruct.MessageContent = strings.TrimSpace(buf.String())
		go func() {
			// Send the actual message here
			sendLogMessage(logStruct)
			ch <- 1
			if logStruct.MessageType == "FATAL" {
				os.Exit(1)
			}
		}()
	}()
}

// writef is applicable for all logs that need to be formatted
func writef(prefixFunc prefixerSignature, log *Logger, clr *color.Color, ch chan int, format string, args ...interface{}) {
	// Create the log that needs to be displayed on stdout
	buf, logStruct := prefixFunc(log)
	clr.Fprintf(buf, format, args...)
	clr.Print(buf.String())
	go func() {
		// Create the log message that needs to be sent to the server, only if the RemoteAvailable flag is set
		if !log.RemoteAvailable {
			ch <- 1
			if logStruct.MessageType == "FATAL" {
				os.Exit(1)
			}
			return
		}
		buf.Reset()
		fmt.Fprintf(buf, format, args...)
		logStruct.MessageContent = strings.TrimSpace(buf.String())
		go func() {
			// Send the actual message here
			sendLogMessage(logStruct)
			ch <- 1
			if logStruct.MessageType == "FATAL" {
				os.Exit(1)
			}
		}()
	}()
}

// writeln is applicable for all logs ending with 'ln'
func writeln(prefixFunc prefixerSignature, log *Logger, clr *color.Color, ch chan int, args ...interface{}) {
	// Create the log that needs to be displayed on stdout
	buf, logStruct := prefixFunc(log)
	clr.Fprint(buf, args...)
	clr.Println(buf.String())
	go func() {
		// Create the log message that needs to be sent to the server, only if the RemoteAvailable flag is set
		if !log.RemoteAvailable {
			ch <- 1
			if logStruct.MessageType == "FATAL" {
				os.Exit(1)
			}
			return
		}
		buf.Reset()
		fmt.Fprintln(buf, args...)
		logStruct.MessageContent = strings.TrimSpace(buf.String())
		go func() {
			// Send the actual message here
			sendLogMessage(logStruct)
			ch <- 1
			if logStruct.MessageType == "FATAL" {
				os.Exit(1)
			}
		}()
	}()
}
