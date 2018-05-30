/*
   Given an array of integers, return indices of the two numbers such that they add up to a
   specific target.
*/

package two_sum

// twoSumI takes an array of integers and a target integer number.
// It returns an array of indices of two numbers that add up to the target.
// Using a hash table method, it takes O(n) run time, O(n) space.
func twoSumI(nums []int, target int) []int {
	tempDict := make(map[int]int)

	for index, value := range nums {
		if _, exists := tempDict[target-value]; exists == false {
			tempDict[value] = index
		} else {
			return []int{tempDict[target-value], index}
		}
	}

	// Return [0, 0] if cannot find any
	return []int{0, 0}
}
