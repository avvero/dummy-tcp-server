// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net"
	"os"
	"bufio"
)

const (
	CONN_HOST = "0.0.0.0"
	CONN_TYPE = "tcp"
)

type SocketService struct {
	tcpPort *string
}

func (c *SocketService) readPump() {
	ln, err := net.Listen(CONN_TYPE, CONN_HOST + ":" + *c.tcpPort)
	if err != nil {
		log.Printf("[socket] Listen error: %v", err)
		os.Exit(1)
	}
	defer ln.Close()
	log.Println("[socket] Tcp server listens on " + CONN_HOST + ":" + *c.tcpPort)
	for {
		log.Println("[socket] Accepting listener")
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("[socket] Accept error: %v", err)
			os.Exit(1)
		}
		log.Printf("[socket] Connection %v established", conn)
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	log.Printf("[socket] Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("[socket] %v", string(netData))
	}
	c.Close()
}