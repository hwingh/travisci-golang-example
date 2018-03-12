package sum

func Sum(x, y int) int {
	return x + y
}

func SumAll(numbers []int) int {
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	return sum
}
