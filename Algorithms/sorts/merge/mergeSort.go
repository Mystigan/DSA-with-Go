package merge

import (
	"errors"
	"strings"
)

// Sort takes a slice of integers, a sorting order and returns a slice sorted using the merge sort algorithm or an error if the sorting order is invalid.
// order can have values: \"asc\" or \"desc\" (irrespective of case).
func Sort(input []int, order string) ([]int, error) {
	if strings.ToLower(order) != "asc" && strings.ToLower(order) != "desc" {
		return nil, errors.New("invalid sorting order")
	}
	n := len(input)
	if n < 2 {
		return input, nil
	}
	left := []int{}
	for i := 0; i < n/2; i++ {
		left = append(left, input[i])
	}
	right := []int{}
	for i := n / 2; i < n; i++ {
		right = append(right, input[i])
	}
	Sort(left, order)
	Sort(right, order)
	return merge(left, right, input, order), nil
}

func merge(left, right, input []int, order string) []int {
	i, j, k, nL, nR := 0, 0, 0, len(left), len(right)
	if strings.ToLower(order) == "asc" {
		for i < nL && j < nR {
			if left[i] < right[j] {
				input[k] = left[i]
				i++
			} else {
				input[k] = right[j]
				j++
			}
			k++
		}
	} else {
		for i < nL && j < nR {
			if left[i] > right[j] {
				input[k] = left[i]
				i++
			} else {
				input[k] = right[j]
				j++
			}
			k++
		}
	}
	for i < nL {
		input[k] = left[i]
		k++
		i++
	}
	for j < nR {
		input[k] = right[j]
		k++
		j++
	}
	return input
}
