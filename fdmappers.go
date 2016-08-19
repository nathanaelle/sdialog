package sdialog // import "github.com/nathanaelle/shesha/sdialog"

import	(
	"os"
	"net"
	"strconv"
	"strings"
	"syscall"
)


type	(
	MapFD interface {
		FDMapper(int) (FileFD,error)
	}

	fd2netListener		struct{}
	fd2netConn		struct{}
	fd2netConnListener	struct{}
)

var (
	NetListenerMapper	MapFD = &fd2netListener{}
	NetConnMapper		MapFD = &fd2netConn{}
	NetConnListenerMapper	MapFD = &fd2netConnListener{}
)





func (_ *fd2netListener)FDMapper(fd int) (conn FileFD, err error) {
	file	:= os.NewFile(uintptr(fd), strings.Join( []string { "@socket_", strconv.Itoa(fd) }, "_" ) )
	t, err	:= net.FileListener(file)
	conn	= t.(FileFD)
	if err	!= nil {
		SD_ALERT.Logf("FDs %d : %v", fd, err)
		syscall.Close(fd)
		return nil,err
	}

	if err	= file.Close(); err != nil {
		SD_ALERT.Logf("FDs %d : %v", fd, err)
		syscall.Close(fd)
		conn.Close()
		return nil,err
	}

	return
}


func (_ *fd2netConn)FDMapper(fd int) (conn FileFD, err error) {
	file	:= os.NewFile(uintptr(fd), strings.Join( []string { "@socket_", strconv.Itoa(fd) }, "_" ) )
	t, err	:= net.FileConn(file)
	conn	= t.(FileFD)
	if err	!= nil {
		SD_ALERT.Logf("FDs %d : %v", fd, err)
		syscall.Close(fd)
		return nil,err
	}

	if err	= file.Close(); err != nil {
		SD_ALERT.Logf("FDs %d : %v", fd, err)
		syscall.Close(fd)
		conn.Close()
		return nil,err
	}

	return
}


func (_ *fd2netConnListener)FDMapper(fd int) (conn FileFD, err error) {
	v, err	:= syscall.GetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_ACCEPTCONN)
	if err != nil {
		return nil,err
	}

	file	:= os.NewFile(uintptr(fd), strings.Join( []string { "@socket_", strconv.Itoa(fd) }, "_" ) )

	switch v {
	case	0:
		t, err	:= net.FileConn(file)
		conn	=  t.(FileFD)
		if err	!= nil {
			SD_ALERT.Logf("FDs %d : %v", fd, err)
			syscall.Close(fd)
			return nil,err
		}

	case	1:
		t, err	:= net.FileListener(file)
		conn	= t.(FileFD)
		if err	!= nil {
			SD_ALERT.Logf("FDs %d : %v", fd, err)
			syscall.Close(fd)
			return nil,err
		}
	}

	if err	= file.Close(); err != nil {
		SD_ALERT.Logf("FDs %d : %v", fd, err)
		syscall.Close(fd)
		conn.Close()
		return nil,err
	}

	return
}
