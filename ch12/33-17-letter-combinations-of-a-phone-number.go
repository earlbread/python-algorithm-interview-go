package ch12

// Recursion
var numLetter = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func makeCombination(result *[]string, digits string, start int, combination []byte) {
	if start >= len(digits) {
		*result = append(*result, string(combination))
		return
	}

	for i := 0; i < len(numLetter[digits[start]]); i++ {
		nextCombination := append(combination, numLetter[digits[start]][i])
		makeCombination(result, digits, start+1, nextCombination)
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	result := []string{}
	makeCombination(&result, digits, 0, []byte{})

	return result
}

// Iteration
func letterCombinations2(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	combinations := []string{""}

	for i := 0; i < len(digits); i++ {
		digit := digits[i]
		letters := numLetter[digit]
		newCombinations := []string{}

		for _, letter := range letters {
			for _, combination := range combinations {
				newCombinations = append(newCombinations, combination+string(letter))
			}
		}

		combinations = newCombinations
	}

	return combinations
}
