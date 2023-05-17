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

func SumAllTails(numbers_range ...[]int) []int {

	var sums []int

	for _, numbers := range numbers_range {

		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
