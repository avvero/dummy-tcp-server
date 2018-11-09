package main

import (
	"flag"
)

import (
	_ "net/http/pprof"
)

type Instance struct {
	Name string
}

type Context struct {
	Instance Instance
	Markers  []string
}

var tcpPort = flag.String("tcpPort", "4562", "tcp server port")

func main() {
	flag.Parse()

	listener := &SocketService{tcpPort: tcpPort}
	listener.readPump()
}
