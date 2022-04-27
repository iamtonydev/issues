package issues

import (
	"errors"
)

func SearchElement(data []int64, target int64) (int64, error) {
	if data == nil {
		return -1, errors.New("input data invalid")
	}

	if data[0] > target || data[len(data)-1] < target {
		return -1, nil
	}

	var start int64 = 0
	end := int64(len(data) - 1)

	for start <= end {
		index := (start + end) / 2
		v := data[index]
		if v == target {
			return index, nil
		} else if v > target {
			end = index - 1
		} else {
			start = index + 1
		}
	}

	return -1, nil
}
