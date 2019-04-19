// +build linux

package sdialog // import "github.com/nathanaelle/sdialog"

import (
	"syscall"
)

type (
	uCred syscall.Ucred
)

func unixCredentials(ucred *uCred) []byte {
	sucred := syscall.Ucred(*ucred)
	return syscall.UnixCredentials(&sucred)
}
