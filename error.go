package	sdialog // import "github.com/nathanaelle/sdialog"

import	(
	"fmt"
	"errors"
)

type	(
	OutOfBoundsLogLevelError	struct {
		Level		LogLevel
		Message		string
	}

	InvalidFDNameError 		struct {
		Name		string
	}

	InvalidStateError		struct {
		State		State
	}
)

var(
	NoSDialogAvailable	error	= errors.New("No SDialog Available")
	NoWatchdogNeeded	error	= errors.New("No Watchdog Needed")
)


func (oob *OutOfBoundsLogLevelError)Error() string {
	return	fmt.Sprintf("invalid LogLevel 0x%02x for message %s\n", oob.Level, oob.Message)
}


func (ifdn *InvalidFDNameError)Error() string {
	return	fmt.Sprintf("invalid name [%s]\n", ifdn.Name)
}


func (isn *InvalidStateError)Error() string {
	return	fmt.Sprintf("state [%v] is invalide\n", isn.State)
}
