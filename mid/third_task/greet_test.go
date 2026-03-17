package main

import "testing"

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
	if err.Error() != "error: --name must not be empty" {
		t.Fatalf("unexpected error message: %q", err.Error())
	}
}

