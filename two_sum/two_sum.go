/*
   Given an array of integers, return indices of the two numbers such that they add up to a
   specific target.
*/

package two_sum

// twoSumI takes an unsorted array of integers and a target integer number.
// It returns an array of indices of two numbers that add up to the target.
// This function uses a hash map approach that maps a value to its index so we can lookup quickly.
// Time complexity: O(n)
// Space complexity: O(n)
func twoSumI(nums []int, target int) []int {
	tempDict := make(map[int]int)

	for index, value := range nums {
		if _, exists := tempDict[target-value]; exists == false {
			tempDict[value] = index
		} else {
			return []int{tempDict[target-value], index}
		}
	}

	// Return [0, 0] if cannot find any.
	return []int{0, 0}
}

// twoSumII takes a sorted array of integers and a target integer number.
// Similar to twoSumII, the goal is to return an array of indices of two numbers that add up to the
// target. The only difference here is that input array is already sorted in ascending order.
// This function uses pointers approach: i start at first index, j start at the last.
// If the sum is larger than the target, decrement j to make it smaller.
// If the sum is smaller than the target, increase i to make it bigger.
// If the sum is equal to the target, we find the answer.
// Time complexity: O(n)
// Space complexity: O(1)
func twoSumII(nums []int, target int) []int {
	sum, i, j := 0, 0, len(nums)-1

	for i < j {
		sum = nums[i] + nums[j]
		if sum == target {
			return []int{i, j}
		}

		if sum < target {
			i++
		} else {
			j--
		}
	}

	// Return [0, 0] if cannot find any.
	return []int{0, 0}
}

// twoSumIIBS takes a sorted array of integers and a target integer number.
// Since all solution above cost us O(n) extra space, we can do better by using binary search.
// For each element x, we could look up the if target - x exists by applying the binary search over
// the sorted array.
// Time complexity: O(n log n)
// Space complexity: O(1)
func twoSumIIBS(nums []int, target int) bool {
	for _, num := range nums {
		if binarySearch(nums, target-num) == true {
			return true
		}
	}
	return false
}
