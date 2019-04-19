package sdialog // import "github.com/nathanaelle/sdialog/v2"

import (
	"bytes"
	"os"
	"path"
	"sync"
	"sync/atomic"
	"time"
)

var lock sync.Locker = &sync.Mutex{}

func initTestingEnv() {
	atomic.StoreInt32(&atomicNoSdAvailable, 0)
	sdcWrite(func(sdc *sdConf) error {
		sdc.logdest = new(bytes.Buffer)
		sdc.notifySocket = "@sdialog_test_notify_socket" + time.Now().Format("2006-01-02T15-04-05.999999999")
		sdc.notifyLocalSocket = "@" + path.Base(os.Args[0]) + "_" + time.Now().Format("15-04-05.999999999")

		return nil
	})
}

func initTestingEnvOut() {
	atomic.StoreInt32(&atomicNoSdAvailable, 0)
	sdcWrite(func(sdc *sdConf) error {
		sdc.logdest = os.Stdout
		sdc.notifySocket = "@sdialog_test_notify_socket" + time.Now().Format("2006-01-02T15-04-05.999999999")
		sdc.notifyLocalSocket = "@" + path.Base(os.Args[0]) + "_" + time.Now().Format("15-04-05.999999999")

		return nil
	})
}

func initTestingEnvNosd() {
	atomic.StoreInt32(&atomicNoSdAvailable, 1)
	sdcWrite(func(sdc *sdConf) error {
		sdc.logdest = new(bytes.Buffer)
		sdc.notifySocket = ""
		sdc.notifyLocalSocket = ""

		return nil
	})
}
