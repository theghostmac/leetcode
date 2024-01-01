package dayone

import "testing"

func Test_lenOfSubString(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		// Basic cases
		{s: "abcabcbb", want: 3},
		{s: "bbbbb", want: 1},
		{s: "pwwkew", want: 3},
		{s: "", want: 0},

		// Substrings with all unique characters.
		{"abcdefg", 7},

		// Edge cases
		{"a", 1},
		{"aa", 1},
		{"aba", 2},
	}

	for _, tc := range tests {
		result := lenOfSubString(tc.s)
		if result != tc.want {
			t.Errorf("lengthOfSubString(%q) = %d, expected %d", tc.s, result, tc.want)
		}
	}
}
