package	sdialog // import "github.com/nathanaelle/sdialog"

import	(
	"fmt"
	"errors"
)

type	(
	outOfBoundsLogLevelError	struct {
		Level		LogLevel
		Message		string
	}

	invalidFDNameError 		struct {
		Name		string
	}

	invalidStateError		struct {
		State		State
	}
)

var(
	NoSDialogAvailable	error	= errors.New("No SDialog Available")
	NoWatchdogNeeded	error	= errors.New("No Watchdog Needed")
)


func (oob *outOfBoundsLogLevelError)Error() string {
	return	fmt.Sprintf("invalid LogLevel 0x%02x for message %s\n", oob.Level, oob.Message)
}


func (ifdn *invalidFDNameError)Error() string {
	return	fmt.Sprintf("invalid name [%s]\n", ifdn.Name)
}


func (isn *invalidStateError)Error() string {
	return	fmt.Sprintf("invalid state [%v]\n", isn.State)
}
