package sdialog // import "github.com/nathanaelle/sdialog"

import (
	"net"
	"os"
)

// Notify send a State to a systemd compatible supervisor
func Notify(states ...State) (err error) {
	var msg, oob []byte
	var notifyConn *net.UnixConn
	var notifySocket string

	if noSdAvailable() {
		return ErrNoSDialogAvailable
	}

	oob = append(oob, unixCredentials(&uCred{
		Pid: int32(os.Getpid()),
		Uid: uint32(os.Getuid()),
		Gid: uint32(os.Getgid()),
	})...)

	sdcWrite(func(sdc *sdConf) error {
		if sdc.notifyConn == nil {
			sdc.notifyConn, err = net.DialUnix("unixgram", &net.UnixAddr{Name: sdc.notifyLocalSocket, Net: "unixgram"}, nil)
			if err != nil {
				LogCRIT.Logf("NOTIFY_SOCKET Error: %s", err.Error())
				return ErrNotifyConnect
			}
		}
		notifyConn = sdc.notifyConn
		notifySocket = sdc.notifySocket

		return nil
	})

	for _, state := range states {
		if !validState(state) {
			LogCRIT.LogError(&invalidStateError{state})
			continue
		}
		m, o := state.State()
		msg = append(msg, m...)
		oob = append(oob, o...)
	}

	_, _, err = notifyConn.WriteMsgUnix(msg, oob, &net.UnixAddr{Name: notifySocket, Net: "unixgram"})
	return
}
