package selection

import (
	"errors"
	"strings"
)

// Sort sorts a provided slice of integers with the specified order using the selection sort algorithm. Returns an error if the order is invalid.
func Sort(nums []int, order string) ([]int, error) {
	var min, index int
	var foundMin bool

	if strings.ToLower(order) == "asc" {
		for i, v := range nums {
			min = v
			foundMin = false
			for j := i; j < len(nums); j++ {
				if nums[j] < min {
					min = nums[j]
					index = j
					foundMin = true
				}
			}
			if foundMin {
				nums[index], nums[i] = nums[i], min
			}
		}
		return nums, nil
	}

	var max int
	var foundMax bool
	if strings.ToLower(order) == "desc" {
		for i, v := range nums {
			max = v
			foundMax = false
			for j := i; j < len(nums); j++ {
				if nums[j] < max {
					max = nums[j]
					index = j
					foundMax = true
				}
			}
			if foundMax {
				nums[index], nums[i] = nums[i], max
			}
		}
		return nums, nil
	}
	return nil, errors.New("invalid order specified. Expected values: \"asc\" or \"desc\"")
}
