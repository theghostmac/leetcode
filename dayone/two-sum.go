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
	// create the empty map for storing
	mapStore := make(map[int]int)

	// now iterate through nums with the first elemnent
	for index, firstElement := range nums {
		// find second element using O(n)
		secondElement := target - firstElement
		if secondIndex, ok := mapStore[secondElement]; ok {
			return []int{secondIndex, index}
		}
		mapStore[firstElement] = index
	}

	return nil
}
