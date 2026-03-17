package main

import (
	"errors"
	"testing"
)

func TestProcessLine_Ping(t *testing.T) {
	resp, closeConn, err := processLine("PING")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if closeConn {
		t.Fatalf("expected closeConn=false")
	}
	if resp != "PONG" {
		t.Fatalf("want %q, got %q", "PONG", resp)
	}
}

func TestProcessLine_Quit(t *testing.T) {
	resp, closeConn, err := processLine("QUIT")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !closeConn {
		t.Fatalf("expected closeConn=true")
	}
	if resp != "BYE" {
		t.Fatalf("want %q, got %q", "BYE", resp)
	}
}

func TestProcessLine_Echo_PreservesCase(t *testing.T) {
	resp, closeConn, err := processLine("ECHO Hi There")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if closeConn {
		t.Fatalf("expected closeConn=false")
	}
	if resp != "ECHO: Hi There" {
		t.Fatalf("want %q, got %q", "ECHO: Hi There", resp)
	}
}

func TestProcessLine_Unknown(t *testing.T) {
	resp, closeConn, err := processLine("NOPE")
	if err == nil {
		t.Fatalf("expected error, got nil (resp=%q)", resp)
	}
	if !errors.Is(err, ErrUnknownCommand) {
		t.Fatalf("expected ErrUnknownCommand, got %v", err)
	}
	if closeConn {
		t.Fatalf("expected closeConn=false")
	}
}

func TestProcessLine_Empty(t *testing.T) {
	resp, closeConn, err := processLine("   ")
	if err == nil {
		t.Fatalf("expected error, got nil (resp=%q)", resp)
	}
	if !errors.Is(err, ErrEmptyLine) {
		t.Fatalf("expected ErrEmptyLine, got %v", err)
	}
	if closeConn {
		t.Fatalf("expected closeConn=false")
	}
	if resp != "ERR empty" {
		t.Fatalf("want %q, got %q", "ERR empty", resp)
	}
}

