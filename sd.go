package	sdialog // import "github.com/nathanaelle/sdialog"


import	(
	"os"
	"fmt"
	"time"
	"strconv"
	"path"
)

var	(
	no_sd_available		bool
)


func init() {
	var err	error

	notify_socket = os.Getenv("NOTIFY_SOCKET")
	os.Unsetenv("NOTIFY_SOCKET")
	if notify_socket == "" {
		no_sd_available = true
		return
	}

	notify_local_socket = "@_"+path.Base(os.Args[0])+"_"+time.Now().Format("15-04-05.999999999")



	s_wdu	:= os.Getenv("WATCHDOG_USEC")
	s_pid	:= os.Getenv("LISTEN_PID")
	s_fds	:= os.Getenv("LISTEN_FDS")

	os.Unsetenv("LISTEN_PID")
	os.Unsetenv("LISTEN_FDS")
	os.Unsetenv("WATCHDOG_USEC")

	if s_wdu != "" {
		watchdog_usec, err = strconv.Atoi(s_wdu)
		if err != nil {
			SD_ALERT.Log(fmt.Sprintf("WATCHDOG_USEC Error: %s", err.Error()))
		}
	}

	if s_pid != "" {
		l_pid,err = strconv.Atoi(s_pid)
		if  err != nil {
			l_pid = 0
			SD_ALERT.Logf("LISTEN_PID : %v", err)
		}
	}

	if s_fds != "" {
		n_fds,err = strconv.Atoi(s_fds)
		if  err != nil {
			n_fds = 0
			SD_ALERT.Logf("LISTEN_FDS : %v", err)
		}
		n_fds	+= sd_fds_start
	}
}
