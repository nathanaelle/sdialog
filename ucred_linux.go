// +build linux

package sdialog // import "github.com/nathanaelle/sdialog"

import	(
	"syscall"
)


type	(
	Ucred	syscall.Ucred
)


func UnixCredentials(ucred *Ucred) []byte {
	sucred	:= syscall.Ucred(*ucred)
	return	syscall.UnixCredentials( &sucred )
}
