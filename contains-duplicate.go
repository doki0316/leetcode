// input: nums = [1, 2, 3, 1]
// return true
package main

import "sort"

func containsDuplicate(nums []int) bool {
	InitNums := len(nums) //4

	for i := 0; i < InitNums; i++ {

		for j := i + 1; j < InitNums; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

// [1, 1, 2, 3]

func containsDuplicate(nums []int) bool {
	InitNums := len(nums)
	sort.Ints(nums)
	for i := 1; i < InitNums; i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}
