package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("ello world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("hello by name", func(t *testing.T) {
		got := Hello("bashbunni", "")
		want := "Hello bashbunni"
		assertCorrectMessage(t, got, want)
	})

	t.Run("name in french", func(t *testing.T) {
		got := Hello("bash", "french")
		want := "Bonjour bash"
		assertCorrectMessage(t, got, want)
	})
}

func TestUrMom(t *testing.T) {
	got := UrMom()
	want := "ur mom"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
