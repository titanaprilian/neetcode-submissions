func hasDuplicate(nums []int) bool {
	numsMap := make(map[int]int)
	for i := range nums {
		numsMap[nums[i]]++

		if numsMap[nums[i]] > 1 {
			return true
			break
		}
	}

	return false
}
