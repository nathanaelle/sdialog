package sdialog // import "github.com/nathanaelle/sdialog"

import (
	"bytes"
	"fmt"
)

type (
	// State describe a systemd compatible state
	State interface {
		State() (msg, oob []byte)
		String() string
	}

	simpleState struct {
		dumbvar []byte
	}

	msgState struct {
		state   string
		message string
	}
)

func validState(s State) bool {
	m, _ := s.State()
	l := len(m)

	if m[l-1] != '\n' {
		return false
	}

	if bytes.IndexByte(m[1:l-2], '=') < 0 {
		return false
	}

	return true
}

func (s *simpleState) State() (msg, oob []byte) {
	msg = []byte(s.dumbvar[:])
	return
}

func (s *simpleState) String() string {
	return string(s.dumbvar[0 : len(s.dumbvar)-1])
}

func (ms *msgState) State() (msg, oob []byte) {
	s := []byte(ms.state)
	m := []byte(ms.message)

	msg = make([]byte, 0, len(s)+len(m)+2)
	msg = append(msg, s...)
	msg = append(msg, '=')
	msg = append(msg, m...)
	msg = append(msg, '\n')

	return
}

func (ms *msgState) String() string {
	r, _ := ms.State()
	return string(r[0 : len(r)-1])
}

// Ready describe the state READY for systemd supervisor
func Ready() State {
	return &simpleState{[]byte("READY=1\n")}
}

// Reloading describe the state RELOADING for systemd supervisor
func Reloading() State {
	return &simpleState{[]byte("RELOADING=1\n")}
}

// Stopping describe the state STOPPING for systemd supervisor
func Stopping() State {
	return &simpleState{[]byte("STOPPING=1\n")}
}

// Status describe the state STATUS for systemd supervisor
func Status(status string) State {
	return &msgState{"STATUS", status}
}

// BusError describe the state BUSSERROR for systemd supervisor
func BusError(status string) State {
	return &msgState{"BUSERROR", status}
}

// MainPid describe the state MAINPID for systemd supervisor
func MainPid(pid int) State {
	return &msgState{"MAINPID", fmt.Sprintf("%d", pid)}
}

// Errno describe the state ERRNO for systemd supervisor
func Errno(errno int) State {
	return &msgState{"ERRNO", fmt.Sprintf("%d", errno)}
}
