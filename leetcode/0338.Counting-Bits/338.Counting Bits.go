package leetcode

// 奇偶数
func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		if i % 2 == 1 {
			res[i] = res[i-1]+1
		} else {
			res[i] = res[i/2]
		}
	}
	return res
}

func countBits2(num int) []int {
	res := make([]int, num+1)
	for i := 0; i <= num; i++ {
		count := 0
		u := i
		for u != 0 {
			// if u & 1 == 1 {
			//     count += 1
			// }
			// u = u >> 1
			u = u & (u - 1)
			count += 1
		}
		res[i] = count
	}
	return res
}