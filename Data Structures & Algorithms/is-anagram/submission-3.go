func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}

	/* 
	
		order string frist 
		time complexcity O(nlogn)
		space comp O(n)
	
	*/

	// r1 := []rune(s)
	// r2 := []rune(t)


	// sort.Slice(r1, func(i, j int) bool{
	// 	return r1[i] < r1[j]
	// })

	// sort.Slice(r2, func(i, j int) bool{
	// 	return r2[i] < r2[j]
	// })

	// // return slices.Equal(r1, r2)

	// return string(r1) == string(r2)


    /* 
		using count letter O(n) and time O(k)
    */

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
