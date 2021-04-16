package leetcode

func searchInsert(nums []int, target int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	if n == 1 {
		if target <= nums[0] {
			return 0
		} else {
			return 1
		}
	}

	left, right := 0, n-1

	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
