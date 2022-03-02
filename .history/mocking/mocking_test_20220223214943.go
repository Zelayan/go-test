package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {

    t.Run("prints 3 to Go!", func(t *testing.T) {
        buffer := &bytes.Buffer{}
        Countdown(buffer, &CountdownOperationsSpy{})

        got := buffer.String()
        want := `3
2
1
Go!`

        if got != want {
            t.Errorf("got '%s' want '%s'", got, want)
        }
    })

    t.Run("sleep after every print", func(t *testing.T) {
        spySleepPrinter := &CountdownOperationsSpy{}
        Countdown(spySleepPrinter, spySleepPrinter)

        want := []string{
            sleep,
            write,
            sleep,
            write,
            sleep,
            write,
            sleep,
            write,
        }

        if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
            t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
        }
    })
}

func TestShellProcessSuccess(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}
	fmt.Fprintf(os.Stdout, testStdoutValue)
	os.Exit(0)
}