package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, n := range nums {
		another := target - n
		if idx, ok := m[another]; ok {
			return []int{i, idx}
		}
		m[n] = i
	}
	return nil
}
