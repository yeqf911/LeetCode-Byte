package leetcode

func removeDuplicates(nums []int) int {
	left, right := 0, 1
	for right < len(nums) {
		if nums[right] == nums[left] {
			right++
		} else {
			nums[left+1] = nums[right]
			right++
			left++
		}
	}
	return left+1
}
