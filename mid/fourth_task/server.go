package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	readTimeout  = 30 * time.Second
	writeTimeout = 5 * time.Second
)

func main() {
	addr := flag.String("addr", "127.0.0.1:9090", "tcp listen address")
	flag.Parse()

	ln, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("listen %s: %v", *addr, err)
	}
	defer ln.Close()

	log.Printf("tcp server listening on %s", ln.Addr().String())

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept: %v", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	remote := conn.RemoteAddr().String()
	log.Printf("connected: %s", remote)
	defer log.Printf("disconnected: %s", remote)

	sc := bufio.NewScanner(conn)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for sc.Scan() {
		_ = conn.SetReadDeadline(time.Now().Add(readTimeout))
		_ = conn.SetWriteDeadline(time.Now().Add(writeTimeout))

		resp, closeConn, _ := processLine(sc.Text())
		if _, err := fmt.Fprintln(conn, resp); err != nil {
			return
		}
		if closeConn {
			return
		}
	}

	if err := sc.Err(); err != nil {
		log.Printf("read %s: %v", remote, err)
	}
}

