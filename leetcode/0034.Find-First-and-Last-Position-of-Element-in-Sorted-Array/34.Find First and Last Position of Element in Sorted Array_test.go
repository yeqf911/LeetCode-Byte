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
	want []int
}

type question struct {
	para
	ans
}

func Test_searchRange(t *testing.T) {
	qs := []question{
		{
			para{nums: []int{1, 2, 3, 3, 4}, target: 3},
			ans{want: []int{2, 3}},
		},
		{
			para{[]int{1}, 1},
			ans{[]int{1}},
		},
	}
	fmt.Println(utils.SplitLine)
	for _, q := range qs {
		p, a := q.para, q.ans
		fmt.Printf(utils.TestLogFormat, q, searchRange(p.nums, p.target), a)
	}
	fmt.Println()
}
