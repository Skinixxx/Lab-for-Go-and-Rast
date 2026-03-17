package main

import (
	"errors"
	"testing"
)

func TestGreet_OK(t *testing.T) {
	got, err := greet("Alice")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "hello, Alice\n"
	if got != want {
		t.Fatalf("want %q, got %q", want, got)
	}
}

func TestGreet_EmptyName(t *testing.T) {
	got, err := greet("")
	if err == nil {
		t.Fatalf("expected error, got nil (output=%q)", got)
	}
	if got != "" {
		t.Fatalf("expected empty output, got %q", got)
	}
	if !errors.Is(err, ErrEmptyName) {
		t.Fatalf("expected ErrEmptyName, got %v", err)
	}
}

