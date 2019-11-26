package main

import (
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
