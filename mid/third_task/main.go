package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "world", "why to greet")
	flag.Parse()
	if *name == "" {
		fmt.Fprintln(os.Stderr, "error: --name must not be empty")
		os.Exit(2)
	}
	fmt.Printf("hello, %s\n", *name)
}
