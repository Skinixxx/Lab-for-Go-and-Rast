package main

import (
	"errors"
	"strings"
)

var (
	ErrEmptyLine      = errors.New("empty line")
	ErrUnknownCommand = errors.New("unknown command")
)

func processLine(line string) (resp string, closeConn bool, err error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return "ERR empty", false, ErrEmptyLine
	}

	upper := strings.ToUpper(line)
	switch {
	case upper == "PING":
		return "PONG", false, nil
	case upper == "QUIT":
		return "BYE", true, nil
	case strings.HasPrefix(upper, "ECHO "):
		// Preserve original case after "ECHO "
		return "ECHO: " + line[5:], false, nil
	default:
		return "ERR unknown command (use PING, ECHO <text>, QUIT)", false, ErrUnknownCommand
	}
}

