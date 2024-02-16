package array

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sumAll []int
	for _, numbers := range numbersToSum {
		sumAll = append(sumAll, Sum(numbers))
	}
	return sumAll
}

func SumAllTail(numberToSum ...[]int) []int {
	var sumAllTail []int
	for _, numbers := range numberToSum {
		if len(numbers) == 0 {
			sumAllTail = append(sumAllTail, 0)
		} else {
			sumAllTail = append(sumAllTail, Sum(numbers[1:]))
		}
	}

	return sumAllTail
}
