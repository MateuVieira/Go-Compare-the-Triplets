package main

import (
	"reflect"
	"testing"
)

// Complete the compareTriplets function below.
func TestCompareTriplets(t *testing.T) {

	t.Run("Base case", func(t *testing.T) {
		a := []int{5, 6, 7}

		b := []int{3, 6, 10}

		got := compareTriplets(a, b)
		want := []int{1, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, a: %v b: %v", got, want, a, b)
		}
	})

	t.Run("Base with 10 itens", func(t *testing.T) {
		a := []int{5, 6, 7, 1, 8, 9, 20, 1, 50, 15}

		b := []int{3, 6, 10, 8, 8, 10, 3, 5, 45, 10}

		got := compareTriplets(a, b)
		want := []int{4, 4}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, a: %v b: %v", got, want, a, b)
		}
	})

	t.Run("Base case with negative int", func(t *testing.T) {
		a := []int{5, -6, 7, 1}

		b := []int{3, 6, 10, -1}

		got := compareTriplets(a, b)
		want := []int{2, 2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, a: %v b: %v", got, want, a, b)
		}
	})
}
