package sdialog // import "github.com/nathanaelle/sdialog/v2"

import (
	"fmt"
	"io"
	"net"
	"os"
	"path"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type (
	sdConf struct {
		lock sync.Locker

		notifyConn        *net.UnixConn
		notifySocket      string
		notifyLocalSocket string

		logdest io.Writer

		watchdogUsec    int
		watchdogRunning bool

		listenPid int
		fdsLen    int
	}
)

var (
	atomicNoSdAvailable int32

	sdc = sdcInit()
)

func noSdAvailable() bool {
	return atomic.LoadInt32(&atomicNoSdAvailable) == 1
}

func init() {
	var err error

	notifySocket := os.Getenv("NOTIFY_SOCKET")
	os.Unsetenv("NOTIFY_SOCKET")
	if notifySocket == "" {
		atomic.StoreInt32(&atomicNoSdAvailable, 1)
		return
	}

	sdcWrite(func(sdc *sdConf) error {
		sdc.notifySocket = notifySocket
		sdc.notifyLocalSocket = "@_" + path.Base(os.Args[0]) + "_" + time.Now().Format("15-04-05.999999999")

		wduStr := os.Getenv("WATCHDOG_USEC")
		pidStr := os.Getenv("LISTEN_PID")
		fdsStr := os.Getenv("LISTEN_FDS")

		os.Unsetenv("LISTEN_PID")
		os.Unsetenv("LISTEN_FDS")
		os.Unsetenv("WATCHDOG_USEC")

		if wduStr != "" {
			sdc.watchdogUsec, err = strconv.Atoi(wduStr)
			if err != nil {
				LogCRIT.Log(fmt.Sprintf("WATCHDOG_USEC Error: %s", err.Error()))
			}
		}

		if pidStr != "" {
			sdc.listenPid, err = strconv.Atoi(pidStr)
			if err != nil {
				sdc.listenPid = 0
				LogCRIT.Logf("LISTEN_PID : %v", err)
			}
		}

		if fdsStr != "" {
			sdc.fdsLen, err = strconv.Atoi(fdsStr)
			if err != nil {
				sdc.fdsLen = 0
				LogCRIT.Logf("LISTEN_FDS : %v", err)
			}
			sdc.fdsLen += fdsStart
		}
		return nil
	})
}

func isMainpid() bool {
	listenPid := 0
	sdcRead(func(sdc sdConf) error {
		listenPid = sdc.listenPid
		return nil
	})
	return (listenPid == 0) || (os.Getpid() != listenPid)
}

func sdcInit() *sdConf {
	return &sdConf{
		lock:    &sync.Mutex{},
		logdest: os.Stderr,
	}
}

func sdcRead(f func(sdConf) error) error {
	sdc.lock.Lock()
	defer sdc.lock.Unlock()

	return f(*sdc)
}

func sdcWrite(f func(*sdConf) error) error {
	sdc.lock.Lock()
	defer sdc.lock.Unlock()

	return f(sdc)
}
