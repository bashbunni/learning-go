package main

func Sum(input []int) int {
	var sum int
	for _, val := range input {
		sum += val
	}
	return sum
}

func SumAll(inputs ...[]int) []int {
	numInput := len(inputs)
	sums := make([]int, numInput)

	for i, val := range inputs {
		sums[i] = Sum(val)
	}
	return sums
}
