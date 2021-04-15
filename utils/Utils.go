package utils

func quicksort(nums []int, left, right int) {
	mid := findMid(nums, left, right)
	quicksort(nums, left, mid-1)
	quicksort(nums, mid+1, right)
}

func findMid(nums []int, left, right int) int {
	pivot := nums[left]
	for left < right {
		for left < right && nums[right] > pivot {
			right--
		}
		nums[left] = nums[right]

		for left < right && nums[left] < pivot {
			left++
		}
		nums[right] = nums[left]
	}

	nums[left] = pivot
	return left
}
