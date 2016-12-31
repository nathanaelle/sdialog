package sdialog // import "github.com/nathanaelle/sdialog"

import	(
	"fmt"
	"unicode"
)


func ValidServiceName(name string) (string, error) {
	if IsSystemService(name) {
		return name, nil
	}

	for i,c := range name {
		switch i {
		case 0:
			if !( unicode.IsNumber(c) || unicode.IsLetter(c) ) {
				return "", fmt.Errorf("Service name invalid [%s] found %x[%s]", name, c, c)
			}
		default:
			if !( unicode.IsNumber(c) || unicode.IsLetter(c) || c == '-' || c == '_' || c == '@' ) {
				return "", fmt.Errorf("Service name invalid [%s] found %x[%s]", name, c, c)
			}
		}
	}

	return name,nil
}


func IsSystemService(name string) bool {
	if name == "_system.start" || name == "_system.stop" {
		return true
	}

	if name == "_system.syslog" || name == "_system.network"  || name == "_system.dev" {
		return true
	}

	return	false
}
