/*
Package sdialog offer a golang way interface to communicate with systemd

the simpliest way to use sdialog is :


	package main

	import	(
		sd "github.com/nathanaelle/sdialog/v2"
	)

	func main() {
		// notify systemd the app is ready
		sd.Notify(sd.Ready())

		sd.LogERR.Log("some message")

		// notify systemd the app is stopping
		sd.Notify(sd.Stopping())
	}





*/
package sdialog // import "github.com/nathanaelle/sdialog/v2"
