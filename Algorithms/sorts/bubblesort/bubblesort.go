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

	if order == "asc" {
		for i := range input {
			swapped := false
			for j := 0; j < len(input)-i-1; j++ {
				if input[j] > input[j+1] {
					input[j], input[j+1] = input[j+1], input[j]
					swapped = true
				}
			}
			if swapped == false {
				break
			}
		}
	} else {
		for i := range input {
			swapped := false
			for j := 0; j < len(input)-i-1; j++ {
				if input[j] < input[j+1] {
					input[j], input[j+1] = input[j+1], input[j]
					swapped = true
				}
			}
			if swapped == false {
				break
			}
		}
	}
	return input, nil
}
