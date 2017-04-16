package ulogger

import (
	"time"
)

// New returns a logger object
func New() *Logger {
	log := new(Logger)
	return log
}

// Info displays a non-critical log message
func (*Logger) Info(args interface{}) {
	infoColor.Print(time.Now().Format("02-01-06 03:04:05"), "\t", args)
}

// Infoln displays a non-critical log message
func (*Logger) Infoln(args interface{}) {
	infoColor.Println(args)
}

// Debug displays a debugging message useful in development environment
func (*Logger) Debug(args interface{}) {
	debugColor.Println(args)
}

// WithFields adds the passed fields and attaches them to the logging object
func (log *Logger) WithFields(fields []DisplayField) {
	for _, field := range fields {
		log.fieldsToDisplays = append(log.fieldsToDisplays, field)
	}
}
