package leetcode

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	diff := 1<<31 - 1
	res := 0
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			currDiff := abs(sum - target)
			if currDiff < diff {
				res, diff = sum, currDiff
			}

			if sum == target {
				return sum
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
