package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)
const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
    Sleep()
}

type SpySleeper struct {
    Calls int
}

func (s *SpySleeper) Sleep() {
    s.Calls++
}

type ConfigurableSleeper struct {
    duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
    time.Sleep(o.duration)
}

type CountdownOperationsSpy struct {
    Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
    s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
    s.Calls = append(s.Calls, write)
    return
}

const write = "write"
const sleep = "sleep"
func Countdown(out io.Writer, sleeper Sleeper) {
    for i := countdownStart; i > 0; i-- {
        sleeper.Sleep()
        fmt.Fprintln(out, i)
    }

    sleeper.Sleep()
    fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
    Countdown(os.Stdout, sleeper)
}

// ShellCommand prints out a fun shell command!
func ShellCommand() (*bytes.Buffer, error) {
	cmd := exec.Command("echo", "'fun!'")

	// Set up byte buffers to read stdout
	var outb bytes.Buffer
	cmd.Stdout = &outb
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return &outb, nil
}