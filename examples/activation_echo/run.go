// +build linux

package main

import (
	"fmt"
	"io"
	"net"

	sd "github.com/nathanaelle/sdialog"
)

func main() {
	if err := sd.LogINFO.Log("ok"); err != nil {
		panic(err)
	}

	ret := sd.FDRetrieve(sd.NetConnMapper)
	if len(ret) != 2 {
		panic(fmt.Errorf("retreived %v sockets !!", len(ret)))
	}
	reader := netconnify(ret[0])
	writer := netconnify(ret[1])
	io.Copy(writer, reader)
}

func netconnify(s sd.FileFD) net.Conn {
	ret, ok := s.(net.Conn)
	if !ok {
		panic(fmt.Errorf("can't netconnify %+v", s))
	}
	return ret
}
