func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	anagramMap := make(map[rune]int)

	for _, char := range s {
		anagramMap[char]++
	}

	for _, char := range t {
		anagramMap[char]--
	}

	for char := range anagramMap {
		if anagramMap[char] != 0 {
			return false
		}
	}

	return true
}
