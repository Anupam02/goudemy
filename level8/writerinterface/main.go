package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("hello from main")
	fmt.Fprintln(os.Stdout, "hello from Fprintf")
	io.WriteString(os.Stdout, "Hello from os.WriteStrig")
}
