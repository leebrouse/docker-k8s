package leetcode

import "sort"

func twoSum1(nums []int, target int) []int {
	sort.Ints(nums)

	var result []int

	left, right := 0, len(nums)-1
	for range nums {
		if (nums[left] + nums[right]) == target {
			result = append(result, left+1, right+1)
			return result
		}

		if (nums[left] + nums[right]) > target {
			right--
		} else {
			left++
		}
	
	}
	
	return nil


}

