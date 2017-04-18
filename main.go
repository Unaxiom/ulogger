package ulogger

import (
	"time"

	"fmt"

	"github.com/fatih/color"
)

// New returns a logger object
func New() *Logger {
	log := new(Logger)
	// Set info colors here
	log.InfoColor = color.New(color.FgHiGreen)
	log.InfoTimeColor = color.New(color.FgHiGreen).Add(color.Underline)
	log.InfoMessageTypeColor = color.New(color.FgHiGreen)
	// Set debug colors here
	log.DebugColor = color.New(color.FgHiWhite).Add(color.BgBlue)
	log.DebugTimeColor = color.New(color.FgHiWhite)
	log.DebugMessageTypeColor = color.New(color.FgHiWhite)
	// Set warning colors here
	log.WarningColor = color.New(color.FgHiYellow)
	log.WarningTimeColor = color.New(color.FgHiYellow)
	log.WarningMessageTypeColor = color.New(color.FgHiYellow)
	// Set error colors here
	log.ErrorColor = color.New(color.FgHiBlue)
	log.ErrorTimeColor = color.New(color.FgHiBlue)
	log.ErrorMessageTypeColor = color.New(color.FgHiBlue)
	// Set critical colors here
	log.CriticalColor = color.New(color.FgHiRed)
	log.CriticalTimeColor = color.New(color.FgHiRed)
	log.CriticalMessageTypeColor = color.New(color.FgHiRed)

	// Set the default log level to info
	log.LogLevel = "info"
	log.SetLogLevel(log.LogLevel)

	return log
}

// SetLogLevel sets the log level of the logger
func (log *Logger) SetLogLevel(level string) {
	log.LogLevel = level
	if log.LogLevel == "debug" {
		log.logLevelCode = 1
	} else if log.LogLevel == "info" {
		log.logLevelCode = 2
	} else if log.LogLevel == "warning" {
		log.logLevelCode = 3
	} else if log.LogLevel == "error" {
		log.logLevelCode = 4
	} else if log.LogLevel == "critical" {
		log.logLevelCode = 5
	}
}

// WithFields adds the passed fields and attaches them to the logging object
func (log *Logger) WithFields(fields []DisplayField) {
	for _, field := range fields {
		log.fieldsToDisplays = append(log.fieldsToDisplays, field)
	}
}

// generateTimestamp returns a logMessage along with the time of creation of this log.
func generateTimestamp() (logMessage, time.Time) {
	var log logMessage
	var timestamp = time.Now()
	log.Timestamp = timestamp.Unix()
	return log, timestamp
}

// sendLogMessage sends the logMessage to the remote URL, if the REMOTE_FLAG is set
func sendLogMessage(log logMessage) {
	// Push the message to the remote URL
	// Also, attach the organization name and the application name here, before composing a new struct
	fmt.Printf("Sending Log Message %#v\n\n", log)
}
