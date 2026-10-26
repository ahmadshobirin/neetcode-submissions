func hasDuplicate(nums []int) bool {

	// n := len(nums)

    // for i := 0; i < n; i++{
	// 	for j := i+1; j < n; j++{
	// 		if nums[i] == nums[j]{
	// 			return true
	// 		}
	// 	}
	// }

	// return false


	seen := make(map[int]bool)

	for _, num := range nums{
		if seen[num]{
			return true
		}
		seen[num] = true
	} 

	return false
}
