package ch17

import "sort"

// Sorting
func isAnagram(s string, t string) bool {
	r1 := []rune(s)
	r2 := []rune(t)

	sort.Slice(r1, func(i, j int) bool {
		return r1[i] < r1[j]
	})
	sort.Slice(r2, func(i, j int) bool {
		return r2[i] < r2[j]
	})

	return string(r1) == string(r2)
}

// Using hashmap
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letters := map[byte]int{}
	for i := 0; i < len(s); i++ {
		letters[s[i]]++
		letters[t[i]]--
	}

	for _, v := range letters {
		if v != 0 {
			return false
		}
	}

	return true
}
