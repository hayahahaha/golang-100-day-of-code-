package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const sleep = "sleep"
const write = "write"

func TestCountDown(t *testing.T) {
	t.Run("Normal run", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleepy := &SpySleepy{}

		CountDown(buffer, spySleepy)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Sleep before run ", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		CountDown(spySleepPrinter, spySleepPrinter)

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

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v ", sleepTime, spyTime.durationSlept)
	}
}

type SpySleepy struct {
	Calls int
}

func (s *SpySleepy) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}