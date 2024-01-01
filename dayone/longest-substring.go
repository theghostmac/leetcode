package dayone

func lenOfSubString(s string) int {
	freqMap := make(map[byte]int)
	previousChar := 0
	maxLength := 0

	// loop over s characters
	for character := 0; character < len(s); character++ {
		// go through the characters and add them to a frequency map.
		characterCount := s[character]
		freqMap[characterCount]++

		// find any duplicated character and shrink the list just before a duplicate.
		for freqMap[characterCount] > 1 {
			prevChar := s[previousChar]
			freqMap[prevChar]--
			previousChar++
		}

		maxLength = findMaxLength(maxLength, character-previousChar+1)
	}

	return maxLength
}

func findMaxLength(a, b int) int {
	if a > b {
		return a
	}
	return b
}
