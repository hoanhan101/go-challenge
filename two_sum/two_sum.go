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
// We could still apply the hash table approach but it costs O(n) extra space, plus it does not
// make use of the fact that the input is already sorted. Therefore, we are using pointers
// approach: i start at first index, j start at the last.
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
