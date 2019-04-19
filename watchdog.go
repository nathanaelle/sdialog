package sdialog // import "github.com/nathanaelle/sdialog"

import (
	"context"
	"sync"
	"time"
)

var (
	watchdogState = &simpleState{[]byte("WATCHDOG=1\n")}
)

// Watchdog run a systemd compatible watchdog
func Watchdog(end context.Context, wg *sync.WaitGroup) error {
	if noSdAvailable() {
		return ErrNoSDialogAvailable
	}

	return sdcWrite(func(sdc *sdConf) error {
		if sdc.watchdogUsec < 1 {
			return ErrNoWatchdogNeeded
		}

		if sdc.watchdogRunning {
			return nil
		}

		// see http://www.freedesktop.org/software/systemd/man/sd_watchdog_enabled.html
		ticker := time.Tick(time.Duration(sdc.watchdogUsec/2) * time.Microsecond)
		sdc.watchdogRunning = true

		if end == nil {
			go watchdogWithoutEnd(ticker, wg)
			return nil
		}

		go watchdogWithEnd(end, ticker, wg)

		return nil
	})
}

func watchdogWithoutEnd(ticker <-chan time.Time, wg *sync.WaitGroup) {
	if wg != nil {
		wg.Add(1)
		defer wg.Done()
	}

	for range ticker {
		Notify(watchdogState)
	}
}

func watchdogWithEnd(end context.Context, ticker <-chan time.Time, wg *sync.WaitGroup) {
	if wg != nil {
		wg.Add(1)
		defer wg.Done()
	}

	for {
		select {
		case <-ticker:
			Notify(watchdogState)
		case <-end.Done():
			return
		}
	}
}
