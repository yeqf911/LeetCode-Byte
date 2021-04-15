package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int
	for i := 0; i < n - 3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n - 2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			L, R := j+1, n-1
			for L < R {
				sum := nums[i] + nums[j] + nums[L] + nums[R]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[L], nums[R]})

					for L < R && nums[L] == nums[L+1] {
						L++
					}

					for L < R && nums[R] == nums[R-1] {
						R--
					}

					L++
					R--
				} else if sum > target {
					R--
				} else {
					L++
				}
			}
		}
	}
	return res
}
