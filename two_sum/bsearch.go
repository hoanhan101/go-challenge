package two_sum

// binarySearch searches over a sorted array of integer and returns a boolean.
func binarySearch(nums []int, target int) bool {
	// Base case
	if len(nums) == 1 {
		if nums[0] == target {
			return true
		} else {
			return false
		}
	}

	// Recursion step
	m := len(nums) / 2

	if nums[m] == target {
		return true
	}

	if nums[m] < target {
		return binarySearch(nums[m:], target)
	} else {
		return binarySearch(nums[:m], target)
	}

	return false
}
