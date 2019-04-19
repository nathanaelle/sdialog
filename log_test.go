package sdialog // import "github.com/nathanaelle/sdialog/v2"

import (
	"bytes"
	"errors"
	"testing"
)

func TestLog(t *testing.T) {
	initTestingEnv()

	tmpErr := errors.New("test error")
	lvl := []LogLevel{LogEMERG, LogALERT, LogCRIT, LogERR, LogWARNING, LogNOTICE, LogINFO, LogDEBUG}

	for i, sdLevel := range lvl {
		Log(sdLevel, "hello")
		sdLevel.Log("world")
		Logf(sdLevel, "foo %d", i+20)
		sdLevel.Logf("bar %d quux", i+30)
		sdLevel.LogError(tmpErr)
	}

	invLvl := []LogLevel{LogLevel('a'), LogLevel(0)}
	for i, sdLevel := range invLvl {
		Log(sdLevel, "hello")
		sdLevel.Log("world")
		Logf(sdLevel, "foo %d", i+20)
		sdLevel.Logf("bar %d quux", i+30)
		sdLevel.LogError(tmpErr)
	}

	expectedOut := `<0>hello
<0>world
<0>foo 20
<0>bar 30 quux
<0>test error
<1>hello
<1>world
<1>foo 21
<1>bar 31 quux
<1>test error
<2>hello
<2>world
<2>foo 22
<2>bar 32 quux
<2>test error
<3>hello
<3>world
<3>foo 23
<3>bar 33 quux
<3>test error
<4>hello
<4>world
<4>foo 24
<4>bar 34 quux
<4>test error
<5>hello
<5>world
<5>foo 25
<5>bar 35 quux
<5>test error
<6>hello
<6>world
<6>foo 26
<6>bar 36 quux
<6>test error
<7>hello
<7>world
<7>foo 27
<7>bar 37 quux
<7>test error
<2>invalid LogLevel 0x61 for message hello
<2>invalid LogLevel 0x61 for message world
<2>invalid LogLevel 0x61 for message foo 20
<2>invalid LogLevel 0x61 for message bar 30 quux
<2>invalid LogLevel 0x61 for message test error
<2>invalid LogLevel 0x00 for message hello
<2>invalid LogLevel 0x00 for message world
<2>invalid LogLevel 0x00 for message foo 21
<2>invalid LogLevel 0x00 for message bar 31 quux
<2>invalid LogLevel 0x00 for message test error
`

	sdcRead(func(sdc sdConf) error {
		testOut := sdc.logdest.(*bytes.Buffer).String()
		if testOut != expectedOut {
			t.Log(testOut)
			t.Error("got wrong output")
		}
		return nil
	})
}

func TestLog_NoSD(t *testing.T) {
	initTestingEnvNosd()

	tmpErr := errors.New("test error")
	lvl := []LogLevel{LogEMERG, LogALERT, LogCRIT, LogERR, LogWARNING, LogNOTICE, LogINFO, LogDEBUG}

	for i, sdLevel := range lvl {
		Log(sdLevel, "hello")
		sdLevel.Log("world")
		Logf(sdLevel, "foo %d", i+20)
		sdLevel.Logf("bar %d quux", i+30)
		sdLevel.LogError(tmpErr)
	}

	expectedOut := ``

	sdcRead(func(sdc sdConf) error {
		testOut := sdc.logdest.(*bytes.Buffer).String()
		if testOut != expectedOut {
			t.Log(testOut)
			t.Error("got wrong output")
		}
		return nil
	})
}
