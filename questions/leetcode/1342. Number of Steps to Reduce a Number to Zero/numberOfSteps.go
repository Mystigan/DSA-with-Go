package leetcode

func numberOfSteps(num int) (count int) {
	for num > 0 {
		if num%2 == 0 {
			num /= 2
			count++
			continue
		}
		num -= 1
		count++
	}
	return
}
