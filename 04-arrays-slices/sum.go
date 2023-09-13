package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums[i] = 0
		} else {
			tail := numbers[1:]
			sums[i] = Sum(tail)
		}
	}

	return sums
}
