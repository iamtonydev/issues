package two_pointers

import "errors"

func TwoSum(data []int, target int) ([]int, error) {
	if data == nil {
		return nil, errors.New("input data invalid")
	}

	left := 0
	right := len(data) - 1

	for left < right {
		if data[left]+data[right] == target {
			return []int{data[left], data[right]}, nil
		} else if data[left]+data[right] < target {
			left++
		} else {
			right--
		}
	}

	return []int{0, 0}, nil
}
