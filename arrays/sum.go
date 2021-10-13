package arrays

// Sum function takes slice of numbers and return their total
func Sum(numbers []int) (total int) {
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTails(numbersToSumTails ...[]int) (tailSums []int) {
	for _, numbers := range numbersToSumTails {
		if len(numbers) > 0 {
			tailSums = append(tailSums, Sum(numbers[1:]))
			continue
		}
		tailSums = append(tailSums, 0)

	}
	return
}
