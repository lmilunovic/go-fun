package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const write = "write"
const sleep = "sleep"

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("sleep before every print", func(t *testing.T) {
		buffer := bytes.Buffer{}
		sleeper := SpySleeper{}

		Countdown(&buffer, &sleeper, 3)

		got := buffer.String()
		want := `3
2
1
Go!`

		if sleeper.Calls != 3 {
			t.Errorf("Sleeper called %d, but wanted 3 times", sleeper.Calls)
		}

		if got != want {
			t.Errorf("got %s, but want %s", got, want)
		}

	})

	t.Run("sleep before every print", func(t *testing.T) {

		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter, 3)

		want := []string{
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
