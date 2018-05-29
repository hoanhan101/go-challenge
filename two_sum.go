/*
   Given an array of integers, return indices of the two numbers such that they add up to a
   specific target. You may assume that each input would have exactly one solution, and
   you may not use the same element twice.
*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Sample test case
	nums := []int{2, 7, 11, 15}
	target := 9

	// Because nums[0] + nums[1] = 2 + 7 = 9, want to return [0, 1].
	expectedResult := []int{0, 1}
	if reflect.DeepEqual(expectedResult, twoSum(nums, target)) == true {
		fmt.Println("Passed")
	}
}

// twoSum takes an array of integers and a target integer number.
// It returns an array of indices of two numbers that add up to the target.
// Using a hash table method, it takes O(n) run time, O(n) space.
func twoSum(nums []int, target int) []int {
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
