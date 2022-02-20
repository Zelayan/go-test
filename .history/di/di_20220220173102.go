package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, name string) {
    fmt.Printf("Hello, %s", name)
}

func main() {
    Greet(os.Stdout, "Elodie")
}