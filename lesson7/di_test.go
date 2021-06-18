// dependency injection test

package main

import (
	"bytes"
	"testing"
)

func TestYo(t *testing.T) {
	buffer := bytes.Buffer{}
	Yo(&buffer, "bashbunni")

	got := buffer.String() // outputs the string rep of buffer value
	want := "Yo, bashbunni"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
