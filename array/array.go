package array

// Sum method
func Sum(array []int) int {
	var sum int
	for _, v := range array {
		sum += v
	}
	return sum
}

// SumAll method
func SumAll(array ...[]int) (sums []int) {
	lengthOfNumbers := len(array)
	sums = make([]int, lengthOfNumbers)

	for i, numbers := range array {
		sums[i] = Sum(numbers)
	}

	return
}
