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
		repeats[v]++
	}

	keys := make([]int, 0, len(repeats))
	for k, _ := range repeats {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return repeats[keys[i]] > repeats[keys[j]]
	})

	var resLen int
	if n < len(keys) {
		resLen = n
	} else {
		resLen = len(keys)
	}

	return keys[:resLen], nil
}
