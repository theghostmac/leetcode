package daytwo

func longestPalindrome(s string) string {
	// iterate through the string.
	// check for potential longest palindromic centers.
	// expand for odd-length and even-length cases.
	// move left and right from the center to find the longest.
	// store entries of the longest found palindromic substring into a variable.
	stringLength := len(s)
	start := 0
	maxLength := 1

	for character := 0; character < stringLength; character++ {
		// case of odd-length center
		leftChar, rightChar := character, character
		// when a palindrome is found
		for leftChar >= 0 && rightChar < stringLength && s[leftChar] == s[rightChar] {
			// if a longer palindrome is found, track it instead of existing one.
			newPalindrome := rightChar - leftChar + 1
			if newPalindrome > maxLength {
				start = leftChar
				maxLength = newPalindrome
			}
		}

		// case of even-length center
		leftChar, rightChar = character, character+1
		for leftChar >= 0 && rightChar < stringLength && s[leftChar] == s[rightChar] {
			newPalindrome := rightChar - leftChar + 1
			if newPalindrome > maxLength {
				start = leftChar
				maxLength = newPalindrome
			}
		}
	}

	return s[start : start+maxLength]
}
