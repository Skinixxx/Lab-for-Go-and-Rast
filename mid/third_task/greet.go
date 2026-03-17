package main

import (
	"errors"
	"fmt"
)

func greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("error: --name must not be empty")
	}
	return fmt.Sprintf("hello, %s\n", name), nil
}

