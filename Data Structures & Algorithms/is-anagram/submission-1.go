func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}

	countletter := make(map[rune]int)

	for _, letter := range s{
		countletter[letter]++
	}

	for _, letter := range t {
		if countletter[letter] == 0{
			return false
		}
		
		countletter[letter]--
	}

	return true
}
