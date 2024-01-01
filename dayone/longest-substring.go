package dayone

func lengthOfLongestSubstring(s string) int {
	// initialize the max length
	// initialize left and right boundaries
	maxLen := 0
	left := 0
	charMap := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		char := s[right]
		charMap[char]++

		for charMap[char] > 1 {
			// Shrink the window from the left until no duplicates
			leftChar := s[left]
			charMap[leftChar]--
			left++
		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
