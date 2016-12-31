// +build linux

package sdialog // import "github.com/nathanaelle/sdialog"

import (
	"bytes"
	"errors"
	"net"
	"os"
	"syscall"
	"testing"
)

func Test_Notify_Ucred(t *testing.T) {
	t.Logf("init\n")
	test_sequence := []State{Ready(), Reloading(), Stopping(), Status("hello"), MainPid(1337)}

	init_testing_env()

	t.Logf("create fake server socket\n")
	srv, err := create_socket_ucred()
	if err != nil {
		t.Error(err)
		return
	}
	defer srv.Close()

	t.Logf("run loop\n")
	go func() {
		for _, s := range test_sequence {
			switch	err := Notify(s); err {
			case	nil:
			case	NoSDialogAvailable:
				t.Error("Env Test isn't detected !!")
			default:
				t.Errorf("Notify loop got : %v", err)
			}
		}
	}()

	for i, s := range test_sequence {
		t.Logf("wait state %d\n", i)
		data := make([]byte, 1<<16)
		oob := make([]byte, 1<<16)
		s_data, s_oob, _, _, err := srv.ReadMsgUnix(data, oob)
		if err != nil {
			t.Error(err)
			return
		}
		data = data[:s_data]
		oob = oob[:s_oob]
		exp_data, exp_oob := s.State()

		if !bytes.Equal(data, exp_data) {
			t.Errorf("data: expected [%v] got [%v]", exp_data, data)
			return
		}

		if len(exp_oob) > s_oob {
			t.Errorf("oob: expected [%v] got [%v]", len(exp_oob), s_oob)
			return
		}

		scms, err := syscall.ParseSocketControlMessage(oob)
		if err != nil {
			t.Errorf("ParseSocketControlMessage : %v", err)
			return
		}

		_, err = syscall.ParseUnixCredentials(&scms[0])
		if err != nil {
			t.Errorf("ParseUnixCredentials : %v", err)
			return
		}

	}
}

func create_socket_ucred() (*net.UnixConn, error) {
	fd, err := syscall.Socket(syscall.AF_UNIX, syscall.SOCK_DGRAM, 0)
	if err != nil {
		return nil, err
	}

	err = syscall.Bind(fd, &syscall.SockaddrUnix{
		Name: notify_socket,
	})
	if err != nil {
		return nil, err
	}

	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_PASSCRED, 1)
	if err != nil {
		return nil, err
	}

	srvFile := os.NewFile(uintptr(fd), notify_socket+"_gofile")
	defer srvFile.Close()

	srv, err := net.FileConn(srvFile)
	if err != nil {
		return nil, err
	}

	dgram_srv, ok := srv.(*net.UnixConn)
	if !ok {
		return nil, errors.New("can't cast dgram_srv")
	}

	return dgram_srv, nil
}
