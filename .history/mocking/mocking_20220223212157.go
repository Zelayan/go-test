package main

import (
	"fmt"
	"io"
	"os"
	"time"
)
const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
    for i := countdownStart; i > 0; i-- {
        time.Sleep(1 * time.Second)
        fmt.Fprintln(out, i)
    }

    time.Sleep(1 * time.Second)
    fmt.Fprint(out, finalWord)
}
func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
    Countdown(os.Stdout, sleeper)
}