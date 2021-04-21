package leetcode

func permute(nums []int) [][]int {
	var traceback func(int, *[]int)
	var res [][]int

	traceback = func(index int, permutation *[]int) {
		if index == len(nums) {
			res = append(res, append([]int{}, *permutation...))
			return
		}
		for i := index; i < len(nums); i++ {
			n := nums[i]
			*permutation = append(*permutation, n)
			traceback(i+1, permutation)
			*permutation = (*permutation)[:len(*permutation)]
		}
	}
	var permutation []int
	traceback(0, &permutation)
	return res
}
