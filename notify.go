package sdialog // import "github.com/nathanaelle/shesha/sdialog"

import	(
	"os"
	"net"
)


var	(
	notifyConn		net.UnixConn
	notify_socket		string
	notify_local_socket	string
)


func Notify(states ...State) (err error) {
	var	msg,oob	[]byte

	if no_sd_available {
		return	NoSDialogAvailable
	}

	oob = append(oob, UnixCredentials(&Ucred{
		Pid:	int32(os.Getpid()),
		Uid:	uint32(os.Getuid()),
		Gid:	uint32(os.Getgid()),
	})...)

	notifyConn, err := net.DialUnix("unixgram", &net.UnixAddr{ Name: notify_local_socket, Net: "unixgram" }, nil)
	if err != nil {
		SD_ALERT.Logf("NOTIFY_SOCKET Error: %s", err.Error())
		return
	}
	defer	notifyConn.Close()

	for _,state := range states {
		if !valid_state(state) {
			SD_ALERT.Error(&InvalidStateError{ state })
			continue
		}
		m,o	:= state.State()
		msg	= append(msg, m...)
		oob	= append(oob, o...)
	}

	_,_,err	= notifyConn.WriteMsgUnix(msg, oob, &net.UnixAddr{ Name: notify_socket, Net: "unixgram" })
	return
}
