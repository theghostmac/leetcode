package dayone

// twoSum uses O(n^2)
func twoSum(nums []int, target int) []int {
	// range through the array first time.
	for i := range nums {
		// range through it the second time starting from i.
		for j := i + 1; j < len(nums); j++ {
			// check if the elements at i and j equals the target.
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	// else return nil.
	return nil
}

// twoSumWithHashmap uses O(n)
func twoSumWithHashmap(nums []int, target int) []int {
	// create empty store map.
	store := make(map[int]int)
	// range through nums
	for index, firstElement := range nums {
		secondElement := target - firstElement
		if jndex, ok := store[secondElement]; ok {
			return []int{jndex, index}
		}

		store[firstElement] = index
	}

	return nil
}

//func twosumbruteforce(nums []int, target int) []int {
//
//}
