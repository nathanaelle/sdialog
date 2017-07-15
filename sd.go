package	sdialog // import "github.com/nathanaelle/sdialog"

import	(
	"io"
	"os"
	"fmt"
	"net"
	"path"
	"sync"
	"time"
	"strconv"
	"sync/atomic"
)

type	(
	sd_conf	struct {
		lock			sync.Locker

		notifyConn		*net.UnixConn
		notify_socket		string
		notify_local_socket	string

		logdest			io.Writer

		watchdog_usec		int
		watchdog_running	bool

		l_pid	int
		n_fds	int
	}
)


var	(
	atomic_no_sd_available	int32	= 0

	sdc	*sd_conf = sdc_init()
)


func no_sd_available() bool {
	return	atomic.LoadInt32(&atomic_no_sd_available) == 1
}

func init() {
	var err	error

	notify_socket := os.Getenv("NOTIFY_SOCKET")
	os.Unsetenv("NOTIFY_SOCKET")
	if notify_socket == "" {
		atomic.StoreInt32(&atomic_no_sd_available, 1)
		return
	}

	sdc_write(func(sdc *sd_conf) error {
		sdc.notify_socket = notify_socket
		sdc.notify_local_socket = "@_"+path.Base(os.Args[0])+"_"+time.Now().Format("15-04-05.999999999")


		s_wdu	:= os.Getenv("WATCHDOG_USEC")
		s_pid	:= os.Getenv("LISTEN_PID")
		s_fds	:= os.Getenv("LISTEN_FDS")

		os.Unsetenv("LISTEN_PID")
		os.Unsetenv("LISTEN_FDS")
		os.Unsetenv("WATCHDOG_USEC")

		if s_wdu != "" {
			sdc.watchdog_usec, err = strconv.Atoi(s_wdu)
			if err != nil {
				SD_CRIT.Log(fmt.Sprintf("WATCHDOG_USEC Error: %s", err.Error()))
			}
		}

		if s_pid != "" {
			sdc.l_pid,err = strconv.Atoi(s_pid)
			if  err != nil {
				sdc.l_pid = 0
				SD_CRIT.Logf("LISTEN_PID : %v", err)
			}
		}

		if s_fds != "" {
			sdc.n_fds,err = strconv.Atoi(s_fds)
			if  err != nil {
				sdc.n_fds = 0
				SD_CRIT.Logf("LISTEN_FDS : %v", err)
			}
			sdc.n_fds	+= sd_fds_start
		}
		return	nil
	})
}

func is_mainpid() bool {
	l_pid	:= 0
	sdc_read(func(sdc sd_conf) error {
		l_pid = sdc.l_pid
		return	nil
	})
	return	(l_pid==0) || (os.Getpid() != l_pid)
}


func sdc_init() *sd_conf  {
	return	&sd_conf {
		lock:		&sync.Mutex{},
		logdest:	os.Stderr,
	}
}

func sdc_read(f func(sd_conf)error)error {
	sdc.lock.Lock()
	defer	sdc.lock.Unlock()

	return f(*sdc)
}

func sdc_write(f func(*sd_conf)error)error {
	sdc.lock.Lock()
	defer	sdc.lock.Unlock()

	return f(sdc)
}
