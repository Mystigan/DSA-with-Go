package array

func findNumbers(nums []int) int {
	numbersWithEvenDigits := 0
	for _, v := range nums {
		if isDigitCountEven(v) {
			numbersWithEvenDigits++
		}
	}
	return numbersWithEvenDigits
}

func isDigitCountEven(value int) bool {
	count := 0
	for value > 0 {
		value /= 10
		count++
	}
	if count%2 == 0 {
		return true
	}
	return false
}
