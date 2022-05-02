package most_frequent_words

import (
	"sort"
	"strings"
)

func MostFrequentWords(text string) ([]string, error) {
	words := strings.Fields(text)
	repeats := make(map[string]int)

	for _, v := range words {
		repeats[v]++
	}

	keys := make([]string, 0, len(repeats))
	for k, _ := range repeats {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		if repeats[keys[i]] == repeats[keys[j]] {
			return keys[i] < keys[j]
		} else {
			return repeats[keys[i]] > repeats[keys[j]]
		}
	})

	var resLen int
	if len(keys) < 10 {
		resLen = len(keys)
	} else {
		resLen = 10
	}

	return keys[:resLen], nil
}
