package sdialog // import "github.com/nathanaelle/sdialog"

import (
	"os"
	"path"
	"time"
	"bytes"
)

func init_testing_env() {
	no_sd_available = false
	logdest = new(bytes.Buffer)
	notify_socket = "@sdialog_test_notify_socket" + time.Now().Format("2006-01-02T15-04-05.999999999")
	notify_local_socket = "@" + path.Base(os.Args[0]) + "_" + time.Now().Format("15-04-05.999999999")
}

func init_testing_env_out() {
	no_sd_available = false
	logdest = os.Stdout
	notify_socket = "@sdialog_test_notify_socket" + time.Now().Format("2006-01-02T15-04-05.999999999")
	notify_local_socket = "@" + path.Base(os.Args[0]) + "_" + time.Now().Format("15-04-05.999999999")
}


func init_testing_env_nosd() {
	no_sd_available = true
	logdest = new(bytes.Buffer)
	notify_socket = ""
	notify_local_socket = ""
}
