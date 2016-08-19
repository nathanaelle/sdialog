package sdialog // import "github.com/nathanaelle/shesha/sdialog"

import (
	"bytes"
	"net"
	"testing"
)

func Test_Notify_NoUcred(t *testing.T) {
	t.Logf("init\n")
	test_sequence := []State{Ready(), Reloading(), Stopping(), Status("hello"), MainPid(1337)}

	init_testing_env()

	t.Logf("create fake server socket\n")
	srv, err := create_socket()
	if err != nil {
		t.Error(err)
		return
	}
	defer srv.Close()

	t.Logf("run loop\n")
	go func() {
		for _, s := range test_sequence {
			err := Notify(s)
			if err != nil {
				panic(err)
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

	}
}

func create_socket() (*net.UnixConn, error) {
	return net.DialUnix("unixgram", &net.UnixAddr{Name: notify_socket, Net: "unixgram"}, nil)
}
