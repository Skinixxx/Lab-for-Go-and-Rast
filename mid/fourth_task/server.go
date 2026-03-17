package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
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
	for sc.Scan() {
		_ = conn.SetWriteDeadline(time.Now().Add(5 * time.Second))

		line := strings.TrimSpace(sc.Text())
		if line == "" {
			if _, err := fmt.Fprintln(conn, "ERR empty"); err != nil {
				return
			}
			continue
		}

		resp, closeConn, err := processLine(line)
		if err != nil {
			if _, werr := fmt.Fprintf(conn, "ERR %s\n", err.Error()); werr != nil {
				return
			}
			if closeConn {
				return
			}
			continue
		}

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

func processLine(line string) (resp string, closeConn bool, err error) {
	upper := strings.ToUpper(line)
	switch {
	case upper == "PING":
		return "PONG", false, nil
	case upper == "QUIT":
		return "BYE", true, nil
	case strings.HasPrefix(upper, "ECHO "):
		// preserve original case after "ECHO "
		return "ECHO: " + line[5:], false, nil
	default:
		return "", false, errors.New("unknown command (use PING, ECHO <text>, QUIT)")
	}
}

