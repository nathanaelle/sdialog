package	sdialog // import "github.com/nathanaelle/shesha/sdialog"


import	(
	"sync"
	"time"
)

var (
	watchdog_usec		int
	watchdog_running	bool
	watchdog_state		= &simple_state{ []byte("WATCHDOG=1\n") }
)


func Watchdog(end <-chan struct{}, wg *sync.WaitGroup)	error {
	if no_sd_available {
		return	NoSDialogAvailable
	}

	if watchdog_usec < 1 {
		return	NoWatchdogNeeded
	}

	if watchdog_running {
		return	nil
	}

	// see http://www.freedesktop.org/software/systemd/man/sd_watchdog_enabled.html
	ticker	:= time.Tick(time.Duration(watchdog_usec/2) * time.Microsecond)
	watchdog_running = true

	if end == nil {
		go func(){
			defer	func(){ watchdog_running=false }()
			if wg != nil {
				wg.Add(1)
				defer wg.Done()
			}

			for {
				select {
				case	<-ticker:
					Notify(watchdog_state)
				}
			}
		}()

		return	nil
	}

	go func(){
		defer	func(){ watchdog_running=false }()
		if wg != nil {
			wg.Add(1)
			defer wg.Done()
		}

		for {
			select {
			case	<-ticker:
				Notify(watchdog_state)
			case	<-end:
				return
			}
		}
	}()

	return	nil
}
