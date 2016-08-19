// +build freebsd netbsd openbsd dragonfly darwin

package sdialog // import "github.com/nathanaelle/shesha/sdialog"



type	(
	Ucred struct {
		Pid int32
		Uid uint32
		Gid uint32
	}
)


func UnixCredentials(ucred *Ucred) []byte {
	SD_WARNING.Log("not implemented in golang")
	return []byte{}
}
