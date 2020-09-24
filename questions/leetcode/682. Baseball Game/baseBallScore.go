package leetcode

import "strconv"

func calPoints(ops []string) int {
	sum, num := 0, 0
	stack := []int{}
	for _, v := range ops {
		switch v {
		case "C":
			sum -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			continue
		case "D":
			num = 2 * stack[len(stack)-1]
			sum += num
		case "+":
			num = stack[len(stack)-1] + stack[len(stack)-2]
			sum += num
		default:
			num, _ = strconv.Atoi(v)
			sum += num
		}
		stack = append(stack, num)
	}
	return sum
}
