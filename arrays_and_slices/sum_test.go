package main

import "testing"
import "reflect"

func assert(t testing.TB, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assert(t, got, want)
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		assert(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Testing multiple collections of any size", func(t *testing.T) {
		first_slice := []int{1, 2}
		second_slice := []int{0, 9}
		got := SumAll(first_slice, second_slice)
		want := []int{3, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %q, want %q", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %q, want %q", got, want)
		}
	}
	t.Run("Testing sum of the tail", func(t *testing.T) {
		first_slice := []int{1, 2}
		second_slice := []int{0, 9}
		got := SumAllTails(first_slice, second_slice)
		want := []int{2, 9}
		checkSums(t, got, want)
	})
	t.Run("sum empty slices", func(t *testing.T) {
		first_slice := []int{}
		second_slice := []int{3, 4, 5}
		got := SumAllTails(first_slice, second_slice)
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
