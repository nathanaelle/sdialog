package sdialog // import "github.com/nathanaelle/sdialog/v2"

import (
	"fmt"
	"log"
)

type (
	// LogLevel describe a log handler for a severity
	LogLevel  byte
	logWriter LogLevel
)

const (
	LogEMERG   LogLevel = iota + '0' // system is unusable
	LogALERT                         // action must be taken immediately
	LogCRIT                          // critical conditions
	LogERR                           // error conditions
	LogWARNING                       // warning conditions
	LogNOTICE                        // normal but significant condition
	LogINFO                          // informational
	LogDEBUG                         // debug-level messages
)

func stderr(l LogLevel, m string) {
	logWriter(l).Write([]byte(m))
}

func (l logWriter) Write(mb []byte) (int, error) {
	encoded := make([]byte, 0, len(mb)+4)
	encoded = append(encoded, '<')
	encoded = append(encoded, byte(l))
	encoded = append(encoded, '>')
	encoded = append(encoded, mb...)
	encoded = append(encoded, '\n')

	sdcRead(func(sdc sdConf) error {
		sdc.logdest.Write(encoded)
		return nil
	})

	return len(mb), nil
}

// Logf log a log to for a LogLevel after formating
func Logf(l LogLevel, format string, v ...interface{}) error {
	return l.Logf(format, v...)
}

// Log log a log to for a LogLevel
func Log(l LogLevel, message string) error {
	return l.Log(message)
}

// Logger expose a log.Logger for a sdialog.LogLevel
func (l LogLevel) Logger(prefix string, flag int) *log.Logger {
	return log.New(logWriter(l), prefix, flag)
}

// Logf log a log to for a LogLevel after formating
func (l LogLevel) Logf(format string, v ...interface{}) error {
	if noSdAvailable() {
		return ErrNoSDialogAvailable
	}

	return l.Log(fmt.Sprintf(format, v...))
}

// Log log a log to for a LogLevel
func (l LogLevel) Log(message string) error {
	if noSdAvailable() {
		return ErrNoSDialogAvailable
	}

	if l < LogEMERG || l > LogDEBUG {
		err := &outOfBoundsLogLevelError{l, message}
		stderr(LogCRIT, err.Error())
		return err
	}

	stderr(l, message)
	return nil
}

// LogError log a error to for a LogLevel
func (l LogLevel) LogError(message error) error {
	if noSdAvailable() {
		return ErrNoSDialogAvailable
	}

	if l < LogEMERG || l > LogDEBUG {
		err := &outOfBoundsLogLevelError{l, message.Error()}
		stderr(LogCRIT, err.Error())
		return err
	}

	stderr(l, message.Error())
	return nil
}
