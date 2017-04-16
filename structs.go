package ulogger

import (
	"github.com/fatih/color"
)

// Logger is the main logging object
type Logger struct {
	fieldsToDisplays []DisplayField
}

// DisplayField stores the name and the value of the field that needs to be printed along with the log message
type DisplayField struct {
	Name  string
	Value interface{}
}

var infoColor = color.New(color.FgHiCyan).Add(color.Underline)
var debugColor = color.New(color.FgGreen).Add(color.BgBlue)
