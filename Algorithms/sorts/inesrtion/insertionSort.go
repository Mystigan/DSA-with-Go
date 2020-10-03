package inesrtion

import (
	"errors"
	"strings"
)

// Takes a slice of integers and sorting order, sorts them using the insertion sort algorithm and returns the sorted slice. Returns an error if invalid sorting order is specified.
// Expected values for sorting order are "asc" or "desc".
func Sort(input []int, order string) ([]int, error) {
	if strings.ToLower(order) == "asc" {
		for i, v := range input {
			index := i
			for index > 0 && input[index-1] > v {
				input[index] = input[index-1]
				index--
			}
			input[index] = v
		}
		return input, nil
	}
	if strings.ToLower(order) == "desc" {
		for i, v := range input {
			index := i
			for index > 0 && input[index-1] < v {
				input[index] = input[index-1]
				index--
			}
			input[index] = v
		}
		return input, nil
	}
	return nil, errors.New("invalid sorting order")
}
