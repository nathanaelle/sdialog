package sdialog // import "github.com/nathanaelle/sdialog"

import	(
	"fmt"
	"log"
)

type	(
	LogLevel	byte
	logWriter	LogLevel

)

const	(
	SD_EMERG	LogLevel = iota+'0'	// system is unusable
	SD_ALERT				// action must be taken immediately
	SD_CRIT					// critical conditions
	SD_ERR					// error conditions
	SD_WARNING				// warning conditions
	SD_NOTICE				// normal but significant condition
	SD_INFO					// informational
	SD_DEBUG				// debug-level messages
)

func stderr(l LogLevel, m string) {
	logWriter(l).Write([]byte(m))
}

func (l logWriter)Write(mb []byte) (int,error)  {
	encoded := make([]byte,0,len(mb)+4)
	encoded = append(encoded, '<')
	encoded = append(encoded, byte(l))
	encoded = append(encoded, '>')
	encoded = append(encoded, mb...)
	encoded = append(encoded, '\n')

	sdc_read(func(sdc sd_conf) error {
		sdc.logdest.Write(encoded)
		return	nil
	})

	return	len(mb),nil
}



func	Logf(l LogLevel, format string, v ...interface{}) error {
	return	l.Logf(format,v...)
}


func	Log(l LogLevel, message string) error {
	return	l.Log(message)
}

func	(l LogLevel)Logger(prefix string, flag int) *log.Logger {
	return	log.New(logWriter(l), prefix, flag)
}


func	(l LogLevel)Logf(format string, v ...interface{}) error {
	if no_sd_available() {
		return	NoSDialogAvailable
	}

	return	l.Log(fmt.Sprintf(format, v...))
}


func (l LogLevel)Log(message string) error {
	if no_sd_available() {
		return	NoSDialogAvailable
	}

	if l < SD_EMERG || l > SD_DEBUG {
		err := &outOfBoundsLogLevelError { l, message }
		stderr(SD_CRIT, err.Error())
		return	err
	}

	stderr(l, message)
	return	nil
}


func (l LogLevel)LogError(message error) error {
	if no_sd_available() {
		return	NoSDialogAvailable
	}

	if l < SD_EMERG || l > SD_DEBUG {
		err := &outOfBoundsLogLevelError { l, message.Error() }
		stderr(SD_CRIT, err.Error())
		return	err
	}

	stderr(l, message.Error())
	return	nil
}
