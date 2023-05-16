package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {

	result := make([]int, len(numbers))

	for i, numSlice := range numbers {
		result[i] = Sum(numSlice)
	}
	return result
}
