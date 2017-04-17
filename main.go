package ulogger

import (
	"bytes"
	"fmt"
	"time"
)

// New returns a logger object
func New() *Logger {
	log := new(Logger)
	return log
}

// Info displays a non-critical log message
func (log *Logger) Info(args interface{}) {
	// infoColor.Print(time.Now().Format("02-01-06 03:04:05"), "\t", args)
	// infoColor.Fprintf()
	infoColor.Print(log.writeToPrint(args))
}

// Infoln displays a non-critical log message
func (*Logger) Infoln(args interface{}) {
	infoColor.Println(args)
}

// Debug displays a debugging message useful in development environment
func (*Logger) Debug(args interface{}) {
	debugColor.Println(args)
}

// writeToPrint writes the logging message to a string and returns it
func (log *Logger) writeToPrint(args interface{}) string {
	buf := new(bytes.Buffer)
	fmt.Fprint(buf, time.Now().Format("02-01-06 03:04:05"), "\t", args)
	return buf.String()
}

// writeToPrintf writes the logging message to a string and returns it
func (log *Logger) writeToPrintf(format string, args ...interface{}) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, format, time.Now().Format("02-01-06 03:04:05"), "\t", args)
	return buf.String()
}

// WithFields adds the passed fields and attaches them to the logging object
func (log *Logger) WithFields(fields []DisplayField) {
	for _, field := range fields {
		log.fieldsToDisplays = append(log.fieldsToDisplays, field)
	}
}
