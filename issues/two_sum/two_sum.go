package two_sum

import "errors"

func TwoSum(data []int, target int) ([]int, error) {
	if data == nil {
		return nil, errors.New("input data invalid")
	}

	previous := make(map[int]struct{})
	res := make([]int, 0, 2)

	for _, v := range data {
		need := target - v
		if _, find := previous[need]; find {
			res = append(res, need, v)
			break
		}

		previous[v] = struct{}{}
	}

	return res, nil
}
