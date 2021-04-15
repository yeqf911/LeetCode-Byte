package leetcode

import (
	"LeetCode-Byte/utils"
	"fmt"
	"testing"
)

type para struct {
	nums []int
}

type ans struct {
	want int
}

type question struct {
	para
	ans
}

func Test_maxArea(t *testing.T) {
	qs := []question{
		{
			para{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			ans{49},
		},
		{
			para{[]int{4, 3, 2, 1, 4}},
			ans{16},
		},
	}

	fmt.Println("------------------------ Leetcode Problem ------------------------")
	for _, q := range qs {
		p, a := q.para, q.ans
		fmt.Printf(utils.TestLogFormat, p, maxArea(p.nums), a)
	}
	fmt.Println()
}
