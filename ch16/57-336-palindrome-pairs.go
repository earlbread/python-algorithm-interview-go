package ch16

type TrieNode struct {
	Children          map[byte]*TrieNode
	WordId            int
	PalindromeWordIds []int
}

type Trie struct {
	root *TrieNode
}

func (trie *Trie) insert(index int, word string) {
	node := trie.root

	for i := 0; i < len(word); i++ {
		c := word[len(word)-i-1]
		if isPalindrome(word[:len(word)-i]) {
			node.PalindromeWordIds = append(node.PalindromeWordIds, index)
		}

		if node.Children[c] == nil {
			node.Children[c] = &TrieNode{
				Children: make(map[byte]*TrieNode),
				WordId:   -1,
			}
		}

		node = node.Children[c]
	}

	node.WordId = index
}

func (trie *Trie) search(index int, word string) [][]int {
	node := trie.root
	result := [][]int{}

	for len(word) > 0 {
		if node.WordId > -1 && isPalindrome(word) {
			result = append(result, []int{index, node.WordId})
		}
		if node.Children[word[0]] == nil {
			return result
		}
		node = node.Children[word[0]]
		word = word[1:]
	}

	if node.WordId > -1 && node.WordId != index {
		result = append(result, []int{index, node.WordId})
	}

	for _, palindromeWordId := range node.PalindromeWordIds {
		result = append(result, []int{index, palindromeWordId})
	}

	return result
}

func isPalindrome(word string) bool {
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}
	return true
}

func palindromePairs(words []string) [][]int {
	trie := Trie{
		&TrieNode{
			Children: make(map[byte]*TrieNode),
			WordId:   -1,
		},
	}

	result := [][]int{}

	for i, w := range words {
		trie.insert(i, w)
	}

	for i, w := range words {
		result = append(result, trie.search(i, w)...)
	}

	return result
}
