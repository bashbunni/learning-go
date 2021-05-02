package integers

import "testing"

func TestHelloWorld(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("sum check", func(t *testing.T) {
		got := sum(2, 2)
		want := 4
		assertCorrectMessage(t, got, want)
	})

	t.Run("product check", func(t *testing.T) {
		got := multiply(2, 2)
		want := 4
		assertCorrectMessage(t, got, want)
	})
}
