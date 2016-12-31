// +build freebsd netbsd openbsd dragonfly darwin

package sdialog // import "github.com/nathanaelle/sdialog"



type	(
	uCred struct {
		Pid int32
		Uid uint32
		Gid uint32
	}
)


func unixCredentials(ucred *uCred) []byte {
	SD_WARNING.Log("not implemented in golang")
	return []byte{}
}
