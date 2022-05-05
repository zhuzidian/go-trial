package arrays

// func Sum(numbers [5]int) int {
// 	var sum int
// 	// for i := 0; i < len(numbers); i++ {
// 	// 	sum += numbers[i]
// 	// }
// 	for _, number := range numbers {
// 		sum += number
// 	}
// 	return sum
// }

func Sum(numbers []int) int {
	var sum int
	// for i := 0; i < len(numbers); i++ {
	// 	sum += numbers[i]
	// }
	for _, number := range numbers {
		sum += number
	}
	return sum
}
