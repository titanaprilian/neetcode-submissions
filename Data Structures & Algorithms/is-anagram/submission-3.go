func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)

	for i := range s {
		sMap[s[i]]++
		tMap[t[i]]++
	}

	for ch := range sMap {
		_, exist := tMap[ch]
		if exist {
			sMap[ch] = sMap[ch] - tMap[ch]
		}
	}

	count := 0
	for ch := range sMap {
		if sMap[ch] > 0 {
			count++
		}
	}

	if count == 0 {
		return count == 0
	}

	return false

}
