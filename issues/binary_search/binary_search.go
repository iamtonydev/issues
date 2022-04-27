package binary_search

import (
	"errors"
)

func SearchElement(data []int, target int) (int, error) {
	if data == nil {
		return -1, errors.New("input data invalid")
	}

	start := 0
	end := len(data) - 1

	for start <= end {
		middle := (start + end) / 2
		if data[middle] == target {
			return middle, nil
		} else if data[middle] > target {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}

	return -1, nil
}
