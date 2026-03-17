package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "world", "who to greet")
	flag.Parse()

	out, err := greet(*name)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	fmt.Print(out)
}
