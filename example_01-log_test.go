package sdialog // import "github.com/nathanaelle/sdialog/v2"

func ExampleLog() {
	initTestingEnvOut()

	lvl := []LogLevel{LogEMERG, LogALERT, LogCRIT, LogERR, LogWARNING, LogNOTICE, LogINFO, LogDEBUG}

	for i, sdLevel := range lvl {
		Log(sdLevel, "hello")
		sdLevel.Log("world")
		Logf(sdLevel, "foo %d", i+20)
		sdLevel.Logf("bar %d quux", i+30)
	}
	// Output:
	// <0>hello
	// <0>world
	// <0>foo 20
	// <0>bar 30 quux
	// <1>hello
	// <1>world
	// <1>foo 21
	// <1>bar 31 quux
	// <2>hello
	// <2>world
	// <2>foo 22
	// <2>bar 32 quux
	// <3>hello
	// <3>world
	// <3>foo 23
	// <3>bar 33 quux
	// <4>hello
	// <4>world
	// <4>foo 24
	// <4>bar 34 quux
	// <5>hello
	// <5>world
	// <5>foo 25
	// <5>bar 35 quux
	// <6>hello
	// <6>world
	// <6>foo 26
	// <6>bar 36 quux
	// <7>hello
	// <7>world
	// <7>foo 27
	// <7>bar 37 quux
}
