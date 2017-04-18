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
	log.InfoTimeColor = color.New(color.FgHiGreen)
	log.InfoMessageTypeColor = color.New(color.FgHiGreen).Add(color.BgHiBlack)
	// Set debug colors here
	log.DebugColor = color.New(color.FgHiWhite)
	log.DebugTimeColor = color.New(color.FgHiWhite)
	log.DebugMessageTypeColor = color.New(color.FgHiWhite).Add(color.BgHiBlack)
	// Set warning colors here
	log.WarningColor = color.New(color.FgHiYellow)
	log.WarningTimeColor = color.New(color.FgHiYellow)
	log.WarningMessageTypeColor = color.New(color.FgHiYellow).Add(color.BgHiBlack)
	// Set error colors here
	log.ErrorColor = color.New(color.FgHiBlue)
	log.ErrorTimeColor = color.New(color.FgHiBlue)
	log.ErrorMessageTypeColor = color.New(color.FgHiBlue).Add(color.BgHiBlack)
	// Set fatal colors here
	log.FatalColor = color.New(color.FgHiRed)
	log.FatalTimeColor = color.New(color.FgHiRed)
	log.FatalMessageTypeColor = color.New(color.FgHiRed).Add(color.BgWhite)

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
	} else if log.LogLevel == "fatal" {
		log.logLevelCode = 5
	} else {
		// Default case is set to 'info'
		log.LogLevel = "info"
		log.SetLogLevel(log.LogLevel)
	}
}

// WithFields adds the passed fields and attaches them to the logging object
func (log *Logger) WithFields(fields []DisplayField) {
	for _, field := range fields {
		log.fieldsToDisplays = append(log.fieldsToDisplays, field)
	}
}

// generateTimestamp returns a logMessage along with the time of creation of this log.
func generateTimestamp(messageType string) (logMessage, time.Time) {
	var log logMessage
	var timestamp = time.Now()
	log.Timestamp = timestamp.Unix()
	log.MessageType = messageType
	return log, timestamp
}

// pushLogMessageToQueue pushes the logMessage to the appropriate queue
func pushLogMessageToQueue(log logMessage) {
	// Also, attach the organization name and the application name here, before composing a new struct
	// Then, acquire the appropriate queue's lock, and push the log message
	fmt.Printf("Sending Log Message %#v\n\n", log)
}

func postLogMessageToServer(log logMessage) {
	// Push the message to the remote URL
	// This function needs should either poll for log messages from the appropriate queues and push to the server
}
