// +build linux

package sdialog // import "github.com/nathanaelle/sdialog/v2"

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
