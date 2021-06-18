package main

import "testing"

// tests

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"toxic1": "get rekd kid"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("toxic1")
		want := "get rekd kid"

		assertError(t, got, want)
	})

	t.Run("unknown", func(t *testing.T) {
		_, err := dictionary.Search("404")
		want := "word not found"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err.Error(), want)
	})
}

// assertions

func assertError(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
