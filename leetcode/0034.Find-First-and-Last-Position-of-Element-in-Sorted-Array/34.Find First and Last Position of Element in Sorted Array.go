package leetcode

func searchRange(nums []int, target int) []int {

	n := len(nums)
	L, R := 0, n-1
	begin, end := -1, -1
	for L <= R {
		mid := (L + R) / 2
		if nums[mid] == target {
			begin, end = mid, mid
			for begin > L && nums[begin-1] == target {
				begin--
			}
			for end < R && nums[end+1] == target {
				end++
			}
			break
		} else if target < nums[mid] {
			R = mid - 1
		} else {
			L = mid + 1
		}
	}
	return []int{begin, end}
}
