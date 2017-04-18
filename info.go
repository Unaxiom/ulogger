package ulogger

import (
	"bytes"
	"fmt"
	"strings"
)

// Info displays a non-critical log message
func (log *Logger) Info(args ...interface{}) {
	// infoColor.Print(time.Now().Format("02-01-06 03:04:05"), "\t", args)
	// infoColor.Fprintf()
	ch := make(chan int)
	go func(ch chan int) {
		buf := new(bytes.Buffer)
		log.InfoColor.Print(log.writeToPrint(buf, args))
		ch <- 1
	}(ch)
	<-ch
}

// Infof displays a non-critical log message according to the format string
func (log *Logger) Infof(format string, args ...interface{}) {
	// Using Print instead of Printf, since the format string would be taken into account in the writeToPrintf method
	ch := make(chan int)
	go func(ch chan int) {
		buf := new(bytes.Buffer)
		log.InfoColor.Print(log.writeToPrintf(buf, format, args...))
		ch <- 1
	}(ch)
	<-ch

}

// Infoln displays a non-critical log message
func (log *Logger) Infoln(args ...interface{}) {
	ch := make(chan int)
	go func(ch chan int) {
		buf, logStruct := infoPrefix(log)
		log.InfoColor.Fprint(buf, args...)
		log.InfoColor.Println(buf.String())
		go func() {
			buf.Reset()
			fmt.Fprintln(buf, args...)
			logStruct.MessageContent = strings.TrimSpace(buf.String())
			go sendLogMessage(logStruct)
			ch <- 1
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
	// fmt.Println("Byte Counter has become ", byteCounter1)
	log.InfoMessageTypeColor.Fprint(buf, logStruct.MessageType)
	fmt.Fprint(buf, " ")

	return buf, logStruct
}
