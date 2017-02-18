package sdialog // import "github.com/nathanaelle/sdialog"

import (
	"bytes"
	"testing"
	"errors"
)

func TestLog(t *testing.T)  {
	init_testing_env()

	t_err	:= errors.New("test error")
	lvl := []LogLevel{SD_EMERG, SD_ALERT, SD_CRIT, SD_ERR, SD_WARNING, SD_NOTICE, SD_INFO, SD_DEBUG}

	for i, sd_level := range lvl {
		Log(sd_level, "hello")
		sd_level.Log("world")
		Logf(sd_level, "foo %d", i+20)
		sd_level.Logf("bar %d quux", i+30)
		sd_level.Error(t_err)
	}

	inv_lvl := []LogLevel{ LogLevel('a'), LogLevel(0) }
	for i, sd_level := range inv_lvl {
		Log(sd_level, "hello")
		sd_level.Log("world")
		Logf(sd_level, "foo %d", i+20)
		sd_level.Logf("bar %d quux", i+30)
		sd_level.Error(t_err)
	}



	expected_out := `<0>hello
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
<1>invalid LogLevel 0x61 for message hello
<1>invalid LogLevel 0x61 for message world
<1>invalid LogLevel 0x61 for message foo 20
<1>invalid LogLevel 0x61 for message bar 30 quux
<1>invalid LogLevel 0x61 for message test error
<1>invalid LogLevel 0x00 for message hello
<1>invalid LogLevel 0x00 for message world
<1>invalid LogLevel 0x00 for message foo 21
<1>invalid LogLevel 0x00 for message bar 31 quux
<1>invalid LogLevel 0x00 for message test error
`

	test_out := logdest.(*bytes.Buffer).String()
	if test_out != expected_out {
		t.Log(test_out)
		t.Error("got wrong output")
	}
}



func TestLog_NoSD(t *testing.T)  {
	init_testing_env_nosd()

	t_err	:= errors.New("test error")
	lvl := []LogLevel{SD_EMERG, SD_ALERT, SD_CRIT, SD_ERR, SD_WARNING, SD_NOTICE, SD_INFO, SD_DEBUG}

	for i, sd_level := range lvl {
		Log(sd_level, "hello")
		sd_level.Log("world")
		Logf(sd_level, "foo %d", i+20)
		sd_level.Logf("bar %d quux", i+30)
		sd_level.Error(t_err)
	}

	expected_out := ``

	test_out := logdest.(*bytes.Buffer).String()
	if test_out != expected_out {
		t.Log(test_out)
		t.Error("got wrong output")
	}
}
