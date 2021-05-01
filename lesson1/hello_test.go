package main

import "testing"

func TestHello(t *testing.T) {
	got := HelloWorld()
	want := "hello world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestUrMom(t *testing.T) {
	got := UrMom()
	want := "hello world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloName(t *testing.T) {
	got := Hello("bashbunni")
	want := "Hello, bashbunni"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
