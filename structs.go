package ulogger

import (
	"github.com/fatih/color"
)

// Logger is the main logging object
type Logger struct {
}

var infoColor = color.New(color.FgHiCyan).Add(color.Underline)
var debugColor = color.New(color.FgGreen).Add(color.BgBlue)
