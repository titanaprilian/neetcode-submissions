func hasDuplicate(nums []int) bool {
	numMaps := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		numMaps[nums[i]]++
		if numMaps[nums[i]] > 1 {
			return true
			break
		}
	}

	return false
}
