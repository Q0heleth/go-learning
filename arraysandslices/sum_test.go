package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		got := Sum(nums)
		expected := 15
		if got != expected {
			t.Errorf("got %d but expected %d, given %v", got, expected, nums)
		}
	})

}
func TestSumTails(t *testing.T) {
	checksum := func(want []int, got []int, t *testing.T) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v but get %v", want, got)
		}
	}
	t.Run("some slices", func(t *testing.T) {
		got := SumTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checksum(want, got, t)
	})
	t.Run("pass empty slice", func(t *testing.T) {
		got := SumTails([]int{}, []int{0, 9})
		want := []int{0, 9}
		checksum(want, got, t)
	})
}
