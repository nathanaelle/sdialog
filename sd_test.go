package sdialog // import "github.com/nathanaelle/sdialog"

import (
	"os"
	"path"
	"sync"
	"time"
	"bytes"
	"sync/atomic"
)

var lock	sync.Locker	= &sync.Mutex{}

func init_testing_env() {
	atomic.StoreInt32(&atomic_no_sd_available, 0)
	sdc_write(func(sdc *sd_conf) error {
		sdc.logdest = new(bytes.Buffer)
		sdc.notify_socket = "@sdialog_test_notify_socket" + time.Now().Format("2006-01-02T15-04-05.999999999")
		sdc.notify_local_socket = "@" + path.Base(os.Args[0]) + "_" + time.Now().Format("15-04-05.999999999")

		return	nil
	})
}

func init_testing_env_out() {
	atomic.StoreInt32(&atomic_no_sd_available, 0)
	sdc_write(func(sdc *sd_conf) error {
		sdc.logdest = os.Stdout
		sdc.notify_socket = "@sdialog_test_notify_socket" + time.Now().Format("2006-01-02T15-04-05.999999999")
		sdc.notify_local_socket = "@" + path.Base(os.Args[0]) + "_" + time.Now().Format("15-04-05.999999999")

		return	nil
	})
}

func init_testing_env_nosd() {
	atomic.StoreInt32(&atomic_no_sd_available, 1)
	sdc_write(func(sdc *sd_conf) error {
		sdc.logdest = new(bytes.Buffer)
		sdc.notify_socket = ""
		sdc.notify_local_socket = ""

		return	nil
	})
}
