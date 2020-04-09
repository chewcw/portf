package main

import (
	"flag"
	"fmt"
	"io"
	"net"
)

var (
	targetServer string
	port         int
)

func init() {
	flag.StringVar(&targetServer, "targetServer", "", "the target (<host>:<port>)")
	flag.IntVar(&port, "port", 7757, "the tunneling port")
}

func main() {
	flag.Parse()

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	proxy, err := net.Dial("tcp", targetServer)
	if err != nil {
		panic(err)
	}

	go copyIO(conn, proxy)
	go copyIO(proxy, conn)
}

func copyIO(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()
	io.Copy(src, dest)
}
