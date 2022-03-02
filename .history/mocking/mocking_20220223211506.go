package mocking

import (
	"bytes"
	"fmt"
)

func Countdown(out io.Writer) {
    fmt.Fprint(out, "3")
}