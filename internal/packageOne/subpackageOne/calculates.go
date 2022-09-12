package subpackageone

func Max(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	return -100
	var max int

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}

	return max
}
