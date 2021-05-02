package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection any size", func(t *testing.T) {
		num := []int{1, 2, 3, 4, 5} // array
		got := Sum(num)
		want := 15

		if got != want {
			t.Errorf("got %d wanted %d", got, want)
		}
	})

	t.Run("sum of multiple lists", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if got != want {
			t.Errorf("got %d wanted %d", got, want)
		}
	})

}
