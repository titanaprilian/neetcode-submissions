func groupAnagrams(strs []string) [][]string {
	gstrs := [][]string{}
	groupIndex := make(map[string]int)

	for _, word := range strs {
		sortedr := []rune(word)	
		sort.Slice(sortedr, func(i, j int) bool {
			return sortedr[i] < sortedr[j]
		})

		canonical := string(sortedr)

		index, exist := groupIndex[canonical]
		if(!exist) {
			index = len(gstrs)	
			groupIndex[canonical] = index
			gstrs = append(gstrs, []string{})
		}

		gstrs[index] = append(gstrs[index], word)
	} 

	return gstrs
}
