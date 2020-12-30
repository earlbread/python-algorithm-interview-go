package ch06

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	for _, str := range strs {
		r := []rune(str)
		sort.Slice(r, func(i, j int) bool {
			return r[i] < r[j]
		})
		sorted := string(r)
		anagrams[sorted] = append(anagrams[sorted], str)
	}

	result := make([][]string, 0)
	for _, strs := range anagrams {
		result = append(result, strs)
	}

	return result
}
