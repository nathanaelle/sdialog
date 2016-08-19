package sdialog // import "github.com/nathanaelle/shesha/sdialog"

import	(
	"fmt"
	"bytes"
)


type	(
	State		interface {
		State()	(msg,oob []byte)
		String() string
	}

	simple_state	struct {
		dumbvar	[]byte
	}

	msg_state	struct {
		state	string
		message string
	}
)


func valid_state(s State) bool {
	m,_	:= s.State()
	l 	:= len(m)

	if m[l-1] != '\n' {
		return	false
	}

	if bytes.IndexByte( m[1:l-2], '=' ) < 0 {
		return	false
	}

	return	true
}



func (s *simple_state)State() (msg,oob []byte) {
	msg	= []byte(s.dumbvar[:])
	return
}


func (s *simple_state)String() string {
	return string(s.dumbvar[0:len(s.dumbvar)-1])
}


func (ms *msg_state)State() (msg,oob []byte) {
	s := []byte(ms.state)
	m := []byte(ms.message)

	msg = make([]byte,0,len(s)+len(m)+2)
	msg = append(msg, s...)
	msg = append(msg, '=')
	msg = append(msg, m...)
	msg = append(msg, '\n')

	return
}


func (s *msg_state)String() string {
	r,_ := s.State()
	return string(r[0:len(r)-1])
}



func Ready() State {
	return	&simple_state{ []byte("READY=1\n") }
}

func Reloading() State {
	return	&simple_state{ []byte("RELOADING=1\n") }
}

func Stopping() State {
	return	&simple_state{ []byte("STOPPING=1\n") }
}

func Status(status string) State {
	return	&msg_state { "STATUS", status }
}

func BusError(status string) State {
	return	&msg_state { "BUSERROR", status }
}

func MainPid(pid int) State {
	return	&msg_state { "MAINPID", fmt.Sprintf("%d", pid) }
}

func Errno(errno int) State {
	return	&msg_state { "ERRNO", fmt.Sprintf("%d", errno) }
}
