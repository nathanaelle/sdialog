package sdialog // import "github.com/nathanaelle/sdialog"

import	(
	"os"
	"net"
)

func Notify(states ...State) (err error) {
	var	msg,oob		[]byte
	var	notifyConn	*net.UnixConn
	var	notify_socket	string

	if no_sd_available() {
		return	NoSDialogAvailable
	}

	oob = append(oob, unixCredentials(&uCred{
		Pid:	int32(os.Getpid()),
		Uid:	uint32(os.Getuid()),
		Gid:	uint32(os.Getgid()),
	})...)


	sdc_write(func(sdc *sd_conf) error {
		if sdc.notifyConn == nil {
			sdc.notifyConn, err = net.DialUnix("unixgram", &net.UnixAddr{ Name: sdc.notify_local_socket, Net: "unixgram" }, nil)
			if err != nil {
				SD_CRIT.Logf("NOTIFY_SOCKET Error: %s", err.Error())
				return NotifyConnectError
			}
		}
		notifyConn	= sdc.notifyConn
		notify_socket	= sdc.notify_socket

		return	nil
	})

	for _,state := range states {
		if !valid_state(state) {
			SD_CRIT.LogError(&invalidStateError{ state })
			continue
		}
		m,o	:= state.State()
		msg	= append(msg, m...)
		oob	= append(oob, o...)
	}

	_,_,err	= notifyConn.WriteMsgUnix(msg, oob, &net.UnixAddr{ Name: notify_socket, Net: "unixgram" })
	return
}
