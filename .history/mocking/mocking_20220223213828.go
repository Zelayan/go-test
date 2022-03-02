package main

import (
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
    // mac平台获取ip地址
    // 执行连续的shell命令时, 需要注意指定执行路径和参数, 否则运行出错
    GetIp()
}

func GetIp() {
	var ip []byte
	var err error
	var cmd *exec.Cmd

	cmd = exec.Command("/bin/sh", "-c", `/sbin/ifconfig en0 | grep -E 'inet ' |  awk '{print $2}'`)
	if ip, err = cmd.Output(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(string(ip))
}