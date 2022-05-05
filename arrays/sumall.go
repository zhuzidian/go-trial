package arrays

func SumAll(numbersToSum ...[]int) []int {
	sums := make([]int, 0)

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
