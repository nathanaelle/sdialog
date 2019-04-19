// +build !linux

package sdialog // import "github.com/nathanaelle/sdialog"

import (
	"bytes"
	"net"
	"testing"
)

func Test_Notify_NoUcred(t *testing.T) {
	t.Logf("init\n")
	testSequence := []State{Ready(), Reloading(), Stopping(), Status("hello"), MainPid(1337)}

	initTestingEnv()

	t.Logf("create fake server socket\n")
	srv, err := createSocket()
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

		if len(oobExp) != oobSize {
			t.Errorf("oob: expected [%v] got [%v]", len(oobExp), oobSize)
			return

		}

	}
}

func Test_Notify_NoUcred_NoSD(t *testing.T) {
	testSequence := []State{Ready(), Reloading(), Stopping(), Status("hello"), MainPid(1337)}
	initTestingEnvNosd()

	for _, s := range testSequence {
		switch err := Notify(s); err {
		case ErrNoSDialogAvailable:
		case nil:
			t.Error("Env Test NoSD isn't detected !!")
		default:
			t.Errorf("Notify loop got : %v", err)
		}
	}
}

func createSocket() (*net.UnixConn, error) {
	notifySocket := ""
	sdcRead(func(sdc sdConf) error {
		notifySocket = sdc.notifySocket
		return nil
	})
	return net.DialUnix("unixgram", &net.UnixAddr{Name: notifySocket, Net: "unixgram"}, nil)
}
