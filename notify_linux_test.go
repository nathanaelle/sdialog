// +build linux

package sdialog

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
	testSequence := []State{Ready(), Reloading(), Stopping(), Status("hello"), MainPid(1337)}

	initTestingEnv()
	ns := ""
	sdcRead(func(sdc sdConf) error {
		ns = sdc.notifySocket
		return nil
	})

	t.Logf("create fake server socket\n")
	srv, err := createSocketUcred(ns)
	if err != nil {
		t.Error(err)
		return
	}
	defer srv.Close()

	t.Logf("run loop\n")
	go func() {
		for _, s := range testSequence {
			switch err := Notify(s); err {
			case nil:
			case ErrNoSDialogAvailable:
				t.Error("Env Test isn't detected !!")
			default:
				t.Errorf("Notify loop got : %v", err)
			}
		}
	}()

	for i, s := range testSequence {
		t.Logf("wait state %d\n", i)
		data := make([]byte, 1<<16)
		oob := make([]byte, 1<<16)
		dataSize, oobSize, _, _, err := srv.ReadMsgUnix(data, oob)
		if err != nil {
			t.Error(err)
			return
		}
		data = data[:dataSize]
		oob = oob[:oobSize]
		dataExp, oobExp := s.State()

		if !bytes.Equal(data, dataExp) {
			t.Errorf("data: expected [%v] got [%v]", dataExp, data)
			return
		}

		if len(oobExp) > oobSize {
			t.Errorf("oob: expected [%v] got [%v]", len(oobExp), oobSize)
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

func createSocketUcred(notifySocket string) (*net.UnixConn, error) {
	fd, err := syscall.Socket(syscall.AF_UNIX, syscall.SOCK_DGRAM, 0)
	if err != nil {
		return nil, err
	}

	err = syscall.Bind(fd, &syscall.SockaddrUnix{
		Name: notifySocket,
	})
	if err != nil {
		return nil, err
	}

	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_PASSCRED, 1)
	if err != nil {
		return nil, err
	}

	srvFile := os.NewFile(uintptr(fd), notifySocket+"_gofile")
	defer srvFile.Close()

	srv, err := net.FileConn(srvFile)
	if err != nil {
		return nil, err
	}

	dgramSrv, ok := srv.(*net.UnixConn)
	if !ok {
		return nil, errors.New("can't cast dgramSrv")
	}

	return dgramSrv, nil
}
