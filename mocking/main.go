package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(writer io.Writer, sleeper Sleeper, times int) {
	for i := times; i >= 1; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, "Go!")
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{}, 5)
}
