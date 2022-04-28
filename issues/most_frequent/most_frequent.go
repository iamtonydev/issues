package most_frequent

import (
	"errors"
	"sort"
)

func MostFrequent(data []int, n int) ([]int, error) {
	if data == nil {
		return nil, errors.New("input data invalid")
	}

	repeats := make(map[int]int)
	for _, v := range data {
		repeats[v] += 1
	}

	kvs := make([][2]int, 0, len(repeats))
	for k, v := range repeats {
		kvs = append(kvs, [2]int{k, v})
	}

	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i][1] > kvs[j][1]
	})

	var resLen int
	if n < len(kvs) {
		resLen = n
	} else {
		resLen = len(kvs)
	}

	res := make([]int, 0, resLen)
	for i := 0; i < resLen; i++ {
		res = append(res, kvs[i][0])
	}

	return res, nil
}
