package sdialog // import "github.com/nathanaelle/sdialog/v2"

import (
	"errors"
	"fmt"
)

type (
	outOfBoundsLogLevelError struct {
		Level   LogLevel
		Message string
	}

	invalidFDNameError struct {
		Name string
	}

	invalidStateError struct {
		State State
	}
)

var (
	// ErrNoSDialogAvailable is rised when systemd (or compatible equivalent) aren't found
	ErrNoSDialogAvailable = errors.New("No SDialog Available")

	// ErrNoWatchdogNeeded is rised when no watchdog is expected
	ErrNoWatchdogNeeded = errors.New("No Watchdog Needed")

	// ErrNotifyConnect is rised when the notification socket isn't available
	ErrNotifyConnect = errors.New("Can't Connect Notify Socket")
)

func (oob *outOfBoundsLogLevelError) Error() string {
	return fmt.Sprintf("invalid LogLevel 0x%02x for message %s", oob.Level, oob.Message)
}

func (ifdn *invalidFDNameError) Error() string {
	return fmt.Sprintf("invalid name [%s]", ifdn.Name)
}

func (isn *invalidStateError) Error() string {
	return fmt.Sprintf("invalid state [%v]\n", isn.State)
}
