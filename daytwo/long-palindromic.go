package daytwo

// given a string s
// return the longest palindromic substring.

func longestPalindrome(s string) string {
	n := len(s)
	start, maxLength := 0, 1

	// iterate through each character of s as potential center of the palindrome.
	for i := 0; i < n; i++ {
		left, right := i, i
		for left >= 0 && right < n && s[left] == s[right] {
			
		}
	}
}
