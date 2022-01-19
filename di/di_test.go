package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Kate")

	got := buffer.String()
	want := "Hello, Kate"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
