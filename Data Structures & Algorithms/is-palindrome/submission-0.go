func isPalindrome(s string) bool {
	reg := regexp.MustCompile(`[^a-zA-Z0-9]`)
	filtered := reg.ReplaceAllString(s, "")
	lowerString := strings.ToLower(filtered)
	charSlice := []rune(lowerString)

	charSliceRev := make([]rune, len(charSlice))

	for i := range charSlice {
		charSliceRev[(len(charSliceRev)-1)-i] = charSlice[i]
	}

	for i := range charSlice {
		if charSlice[i] != charSliceRev[i] {
			return false
		}
	}

	return true
}
