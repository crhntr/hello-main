package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "Hello, out!")
	fmt.Fprintln(os.Stderr, "Hello, err!")
}
