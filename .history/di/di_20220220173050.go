package di

import (
	"bytes"
	"fmt"
	"os"
)

func Greet(writer io.W, name string) {
    fmt.Printf("Hello, %s", name)
}

func main() {
    Greet(os.Stdout, "Elodie")
}