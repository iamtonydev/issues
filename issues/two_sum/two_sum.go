package two_sum

import "errors"

func TwoSum(data []int, target int) ([]int, error) {
	if data == nil {
		return nil, errors.New("input data invalid")
	}

	previous := make(map[int]struct{})

	for _, v := range data {
		diff := target - v
		if _, find := previous[diff]; find {
			return []int{diff, v}, nil
			break
		}

		previous[v] = struct{}{}
	}

	return []int{0, 0}, nil
}
