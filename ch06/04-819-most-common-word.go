package ch06

import (
	"regexp"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	p := strings.ToLower(paragraph)
	r, _ := regexp.Compile("[^a-z ]+")
	p = r.ReplaceAllString(p, " ")
	words := strings.Fields(p)

	bannedMap := make(map[string]bool)
	for _, word := range banned {
		bannedMap[word] = true
	}

	mostCommon := ""
	maxCount := 0
	wordCount := make(map[string]int)
	for _, word := range words {
		if !bannedMap[word] {
			wordCount[word] += 1
		}

		if wordCount[word] > maxCount {
			maxCount = wordCount[word]
			mostCommon = word
		}
	}

	return mostCommon
}
