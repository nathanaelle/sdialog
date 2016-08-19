package sdialog // import "github.com/nathanaelle/shesha/sdialog"

import	(
	"os"
	"fmt"
	"net"
	"strings"
	"syscall"
)

type	(
	FileFD	interface {
		File() (*os.File,error)
		Close() error
		Addr() net.Addr
	}

	FDs	struct {
		name	string
		fds	[]int
	}
)

var (
	l_pid	int
	n_fds	int
)

const	sd_fds_start	int	= 3


func (sd *FDs)String() string {
	return fmt.Sprintf("[%s] = %d fds", sd.name, len(sd.fds))
}


func (sd *FDs)State() (msg,oob []byte) {
	oob	= syscall.UnixRights(sd.fds...)
	msg	= []byte("FDSTORE=1\n")

	if sd.name != "" {
		msg = append(msg, []byte("FDNAME=")...)
		msg = append(msg, []byte(sd.name)...)
		msg = append(msg, '\n')
	}

	return
}


func valid_fdname(name string) bool {
	for _,r := range name {
		switch {
		case	r == ':':
			return	false

		case	r >= 127:
			return	false

		case	r < ' ':
			return	false
		}
	}
	return	true
}



func FDStore(name string, ifaces ...FileFD) State {
	if !valid_fdname(name) {
		SD_ALERT.Logf( "%v", &InvalidFDNameError { name } )
		return	nil
	}

	s := &FDs {
		name: name,
	}

	for _,listener := range ifaces {
		fd,err := listener.File()
		if  err != nil {
			SD_ALERT.Logf("%v : %v", listener.Addr, err)
			continue
		}
		s.fds = append(s.fds, int(fd.Fd()))
	}

	return	s
}


func FDRetrieve(mapper MapFD) (ret []FileFD) {
	if no_sd_available {
		return
	}

	if os.Getpid() != l_pid {
		SD_ALERT.Logf("LISTEN_PID : expected %d got %d", os.Getpid(), l_pid)
		return
	}

	if n_fds <= sd_fds_start {
		return
	}

	if mapper == nil {
		mapper = &fd2netConnListener{}
	}

	for fd	:= sd_fds_start ; fd < n_fds ; fd++ {
		syscall.CloseOnExec(fd)
		l, err	:= mapper.FDMapper(fd)
		if err	!= nil {
			SD_ALERT.Logf("FDs %d : %v", fd, err)
			continue
		}
		ret = append(ret, l)
	}
	return
}


func FDRetrieveByName(mapper MapFD) (ret map[string][]FileFD) {
	lili	:= FDRetrieve(mapper)
	if len(lili) == 0 {
		return
	}

	ret	=  make(map[string][]FileFD)
	l_fdn	:= strings.Split(os.Getenv("LISTEN_FDNAMES"), ":")
	if len(l_fdn) == 0 {
		ret["unknown"] = lili
		return
	}

	if len(l_fdn) != len(lili) {
		SD_ALERT.Logf("FD %v : %v", len(l_fdn), len(lili))
		return
	}

	for i,_ := range l_fdn {
		name := l_fdn[i]
		if !valid_fdname(name) {
			name = "unknown"
		}
		if name == "" {
			name = "unknown"
		}
		switch _,ok := ret[name]; ok {
		case	false:
			ret[name] = []FileFD { lili[i] }
		case	true:
			ret[name] = append(ret[name], lili[i])
		}
	}

	return
}
