package dayone

func maxSubString(s string) int {
	storeMaxCalculatedLength := 0       // store max length of substring counted so far.
	left := 0                           // left boundary of the sliding window.
	right := 0                          // right boundary of the sliding window.
	storeCharsMap := make(map[byte]int) // store the characters and the frequency in the window.

}
