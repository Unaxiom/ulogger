package ulogger

import (
	"github.com/fatih/color"
)

// timeFormat describes the output timestamp format
var timeFormat = "02-01-06 03:04:05"

// remoteFlagString is the parameter that needs to be set during compilation that enables/disables remote logging. Default is "true"
var remoteFlagString = "true"
var remoteFlag bool

var remoteURL = "https://logging.unaxiom.com/newlogmessage"

// Logger is the main logging object
type Logger struct {
	fieldsToDisplays []DisplayField
	// Customizable colors
	// Info colors
	InfoColor            *color.Color // Color of the info message
	InfoTimeColor        *color.Color // Color of the info timestamp
	InfoMessageTypeColor *color.Color // Color of the message type

	// Debug colors
	DebugColor            *color.Color // Color of the debug message
	DebugTimeColor        *color.Color // Color of the debug timestamp
	DebugMessageTypeColor *color.Color // Color of the message type

	// Warning colors
	WarningColor            *color.Color // Color of the warning message
	WarningTimeColor        *color.Color // Color of the warning timestamp
	WarningMessageTypeColor *color.Color // Color of the message type

	// Error colors
	ErrorColor            *color.Color // Color of the error message
	ErrorTimeColor        *color.Color // Color of the error timestamp
	ErrorMessageTypeColor *color.Color // Color of the message type

	// Critical colors
	CriticalColor            *color.Color // Color of the critical message
	CriticalTimeColor        *color.Color // Color of the critical timestamp
	CriticalMessageTypeColor *color.Color // Color of the message type
}

// DisplayField stores the name and the value of the field that needs to be printed along with the log message
type DisplayField struct {
	Name  string
	Value interface{}
}

// logMessage is the internal struct that is posted to the remote log server
type logMessage struct {
	MessageType    string `json:"message_type"`
	Timestamp      int64  `json:"timestamp"`
	MessageContent string `json:"message_content"`
}
