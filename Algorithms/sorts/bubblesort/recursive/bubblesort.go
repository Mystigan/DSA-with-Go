package bubblesort

import (
	"errors"
)

// BubbleSort takes a slice of integers as input and a string flag to decide sorting order and returns the sorted slice.
// order flag is used to determine the sorting order. Accepted values:
//  - "asc": sorts slice ins ascending order.
//  - "desc": sorts slice ins descending order.
//  - Any other value passed returns an error.
func BubbleSort(input []int, order string) ([]int, error) {
	if order != "asc" && order != "desc" {
		return nil, errors.New("invalid ordering. Value must be either of \"asc\" or \"desc\"")
	}
	if len(input) == 0 {
		return nil, errors.New("slice is empty")
	}
	n := len(input)
	if order == "asc" {
		sort(input, n, "asc")
	} else {
		sort(input, n, "desc")
	}
	return input, nil
}

func sort(input []int, n int, order string) {
	if n == 1 {
		return
	}
	if order == "asc" {
		for i := 0; i < n-1; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
			}
		}
	} else {
		for i := 0; i < n-1; i++ {
			if input[i] < input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
			}
		}
	}
	sort(input, n-1, order)
}
