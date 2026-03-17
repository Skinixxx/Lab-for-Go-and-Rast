package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "world", "who to greet")
	flag.Parse()

	out, err := greet(*name)
	if err != nil {
		if errors.Is(err, ErrEmptyName) {
			fmt.Fprintln(os.Stderr, "error: --name must not be empty")
		} else {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		os.Exit(2)
	}

	fmt.Print(out)
}
