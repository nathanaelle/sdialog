package sdialog // import "github.com/nathanaelle/sdialog/v2"

import (
	"fmt"
	"os"
	"strings"
	"syscall"
)

type (
	// FileFD is the minimal common part between all go file descriptor
	FileFD interface {
		File() (*os.File, error)
		Close() error
	}

	fds struct {
		name string
		fds  []int
	}
)

const fdsStart int = 3

func (sd *fds) String() string {
	return fmt.Sprintf("[%s] = %d fds", sd.name, len(sd.fds))
}

func (sd *fds) State() (msg, oob []byte) {
	oob = syscall.UnixRights(sd.fds...)
	msg = []byte("FDSTORE=1\n")

	if sd.name != "" {
		msg = append(msg, []byte("FDNAME=")...)
		msg = append(msg, []byte(sd.name)...)
		msg = append(msg, '\n')
	}

	return
}

func validFdName(name string) bool {
	for _, r := range name {
		switch {
		case r == ':':
			return false

		case r >= 127:
			return false

		case r < ' ':
			return false
		}
	}
	return true
}

// FDStore ask to systemd like supervisor to store an FileFD
func FDStore(name string, ifaces ...FileFD) State {
	if !validFdName(name) {
		LogCRIT.Logf("%v", &invalidFDNameError{name})
		return nil
	}

	s := &fds{
		name: name,
	}

	for id, listener := range ifaces {
		fd, err := listener.File()
		if err != nil {
			LogCRIT.Logf("socket %v : %v", id, err)
			continue
		}
		s.fds = append(s.fds, int(fd.Fd()))
	}

	return s
}

// FDRetrieve ask to systemd like supervisor to retrieve an FileFD
func FDRetrieve(mapper MapFD) (ret []FileFD) {
	if noSdAvailable() {
		return
	}

	if !isMainpid() {
		return
	}

	fdsLen := 0
	sdcRead(func(sdc sdConf) error {
		fdsLen = sdc.fdsLen
		return nil
	})

	if fdsLen <= fdsStart {
		return
	}

	if mapper == nil {
		mapper = &fd2netConnListener{}
	}

	for fd := fdsStart; fd < fdsLen; fd++ {
		syscall.CloseOnExec(fd)
		l, err := mapper.FDMapper(fd)
		if err != nil {
			LogALERT.Logf("FDs %d : %v", fd, err)
			continue
		}
		ret = append(ret, l)
	}
	return
}

// FDRetrieveByName ask to systemd like supervisor to retrieve an named FileFD
func FDRetrieveByName(mapper MapFD) (ret map[string][]FileFD) {
	lili := FDRetrieve(mapper)
	if len(lili) == 0 {
		return
	}

	ret = make(map[string][]FileFD)
	fdnList := strings.Split(os.Getenv("LISTEN_FDNAMES"), ":")
	if len(fdnList) == 0 {
		ret["unknown"] = lili
		return
	}

	if len(fdnList) != len(lili) {
		LogALERT.Logf("FD %v : %v", len(fdnList), len(lili))
		return
	}

	for i := range fdnList {
		name := fdnList[i]
		if !validFdName(name) {
			name = "unknown"
		}
		if name == "" {
			name = "unknown"
		}
		switch _, ok := ret[name]; ok {
		case false:
			ret[name] = []FileFD{lili[i]}
		case true:
			ret[name] = append(ret[name], lili[i])
		}
	}

	return
}
