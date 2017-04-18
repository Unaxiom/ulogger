package ulogger

import "bytes"

// Debug displays a debugging message useful in development environment
func (log *Logger) Debug(args ...interface{}) {
	go func() {
		buf := new(bytes.Buffer)
		log.DebugColor.Print(log.writeToPrint(buf, args...))
	}()

}

// Debugf displays a debugging message
func (log *Logger) Debugf(format string, args ...interface{}) {
	go func() {
		buf := new(bytes.Buffer)
		log.DebugColor.Print(log.writeToPrintf(buf, format, args...))
	}()

}
