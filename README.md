# ulogger

A custom logger to be used for all projects at Unaxiom.

# Usage
```
import "ulogger"

func main() {
    log := ulogger.New()
    log.SetLogLevel("debug") // Possible values are "debug"(least), "info", "warning", "error", "fatal" (highest), in ascending order, with "info being the default"
    log.RemoteAvailable = true // Defines whether logs need to the be streamed to the remote URL
    log.ApplicationName = "Temp Debugger" // Sets the applicaton name
    log.OrganizationName = "New org" // Sets the organization name that this build is licensed to
}
```

# Possible log levels
```
# Debug level logs
log.Debug(args ...interface{})
log.Debugf(format string, args ...interface{})
log.Debugln(args ...interface{})

# Info level logs
log.Info(args ...interface{})
log.Infof(format string, args ...interface{})
log.Infoln(args ...interface{})

# Warning level logs
log.Warning(args ...interface{})
log.Warningf(format string, args ...interface{})
log.Warningln(args ...interface{})

# Error level logs
log.Error(args ...interface{})
log.Errorf(format string, args ...interface{})
log.Errorln(args ...interface{})

# Fatal level logs
log.Fatal(args ...interface{})
log.Fatalf(format string, args ...interface{})
log.Fatalln(args ...interface{})
```

# Modification of output colors
```
log.InfoColor
log.InfoTimeColor
log.InfoMessageTypeColor
log.DebugColor
log.DebugTimeColor
log.DebugMessageTypeColor
log.WarningColor
log.WarningTimeColor
log.WarningMessageTypeColor
log.ErrorColor
log.ErrorTimeColor
log.ErrorMessageTypeColor
log.FatalColor
log.FatalTimeColor
log.FatalMessageTypeColor
```
Each of these values could be assigned an color from the package `github.com/fatih/color`.