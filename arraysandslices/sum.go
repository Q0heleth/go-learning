package main

func Sum(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}
func SumTails(numberstosum ...[]int) []int {
	var sums []int
	for _, slice := range numberstosum {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			tail := slice[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
