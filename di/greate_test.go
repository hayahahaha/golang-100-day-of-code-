package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Bob")

	got := buffer.String()
	want := "Hello Bob"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
