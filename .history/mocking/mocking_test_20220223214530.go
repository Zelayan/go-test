package main

import (
	"bytes"
	"os/exec"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestGetIp(t *testing.T) {
	

}

type RealCommander struct{}

func (c RealCommander) CombinedOutput(command string, args ...string) ([]byte, error) {
	return exec.Command(command, args...).CombinedOutput()
}
const testHash = "3a9a4f7b8a8e1a62691cb3715769f03972fe5597"

type TestCommander struct{}

func (c TestCommander) CombinedOutput(command string, args ...string) ([]byte, error) {
	cs := []string{"-test.run=TestOutput", "--"}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_TEST_OUTPUT=1"}
	out, err := cmd.CombinedOutput()
	return out, err
}

func TestOutput(*testing.T) {
	if os.Getenv("GO_WANT_TEST_OUTPUT") != "1" {
		return
	}

	defer os.Exit(0)
	fmt.Printf(testHash)
}

func TestGetHeadHash(t *testing.T) {
	commander = TestCommander{}
	out, _ := GetHeadHash()
	if string(out) != testHash {
		t.Errorf("Wanted %s, got %s", testHash, string(out))
	}
}