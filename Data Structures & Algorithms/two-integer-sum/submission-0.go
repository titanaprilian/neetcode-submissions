func twoSum(nums []int, target int) []int {
	numMaps := make(map[int]int)

	for i, n := range nums {
		difference := target - nums[i]
		idx, exists := numMaps[difference]

		if exists {
			return []int{idx, i}
		}

		numMaps[n] = i
	}

	return []int{0, 0}

}
