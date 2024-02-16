package array

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collect of 5 number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d number %v", got, want, numbers)
		}
	})
	t.Run("Collect of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d number %v", got, want, numbers)
		}

	})
}

func TestSumAll(t *testing.T) {
	t.Run("testing ", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{1, 2})
		want := []int{6, 3}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v ", got, want)
		}
	})
}

func TestSumAllTail(t *testing.T) {
	checkSum := func(t testing.TB, got, want []int) {
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v ", got, want)
		}
	}

	t.Run("Testing summ all tail", func(t *testing.T) {
		got := SumAllTail([]int{1, 2}, []int{2, 9})
		want := []int{2, 9}
		checkSum(t, got, want)
	})
	t.Run("test sum all tail with 0 length", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{2, 9})
		want := []int{0, 9}
		checkSum(t, got, want)
	})
}
