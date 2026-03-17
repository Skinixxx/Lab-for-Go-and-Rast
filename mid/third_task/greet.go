package main

import (
	"errors"
	"fmt"
)

var ErrEmptyName = errors.New("name must not be empty")

func greet(name string) (string, error) {
	if name == "" {
		return "", ErrEmptyName
	}
	return fmt.Sprintf("hello, %s\n", name), nil
}

