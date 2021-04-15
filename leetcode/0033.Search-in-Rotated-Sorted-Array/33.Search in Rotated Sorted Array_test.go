package leetcode

import (
	"LeetCode-Byte/utils"
	"fmt"
	"testing"
)

type para struct {
	nums   []int
	target int
}

type ans struct {
	want int
}

type question struct {
	para
	ans
}

func Test_search(t *testing.T) {
	qs := []question{
		{
			para{[]int{1, 3}, 3},
			ans{1},
		},
		{
			para{[]int{3, 1}, 1},
			ans{1},
		},
		{
			para{[]int{4, 5, 6, 7, 0, 1, 2}, 3},
			ans{-1},
		},
	}

	fmt.Println(utils.SplitLine)
	for _, q := range qs {
		p, a := q.para, q.ans
		fmt.Printf(utils.TestLogFormat, p, search(p.nums, p.target), a)
	}
	fmt.Println()
}
