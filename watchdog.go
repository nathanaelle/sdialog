package	sdialog // import "github.com/nathanaelle/sdialog"


import	(
	"sync"
	"time"
	"context"
)

var (
	watchdog_state		= &simple_state{ []byte("WATCHDOG=1\n") }
)


func Watchdog(end context.Context, wg *sync.WaitGroup)	error {
	if no_sd_available() {
		return	NoSDialogAvailable
	}

	return	sdc_write(func(sdc *sd_conf) error {
		if sdc.watchdog_usec < 1 {
			return	NoWatchdogNeeded
		}

		if sdc.watchdog_running {
			return	nil
		}

		// see http://www.freedesktop.org/software/systemd/man/sd_watchdog_enabled.html
		ticker	:= time.Tick(time.Duration(sdc.watchdog_usec/2) * time.Microsecond)
		sdc.watchdog_running = true

		if end == nil {
			go watchdog_without_end(ticker, wg)
			return	nil
		}

		go watchdog_with_end(ticker, end, wg)

		return	nil
	})
}


func watchdog_without_end(ticker <-chan time.Time, wg *sync.WaitGroup){
	if wg != nil {
		wg.Add(1)
		defer wg.Done()
	}

	for range ticker {
		Notify(watchdog_state)
	}
}


func watchdog_with_end(ticker <-chan time.Time, end context.Context, wg *sync.WaitGroup){
	if wg != nil {
		wg.Add(1)
		defer wg.Done()
	}

	for {
		select {
		case	<-ticker:
			Notify(watchdog_state)
		case	<-end.Done():
			return
		}
	}
}
