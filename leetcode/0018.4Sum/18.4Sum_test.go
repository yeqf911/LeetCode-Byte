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
	want [][]int
}

type question struct {
	para
	ans
}

func Test_fourSum(t *testing.T) {
	qs := []question{
		{
			para{nums: []int{1, 0, -1, 0, -2, 2}, target: 0},
			ans{[][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		},
	}
	fmt.Println(utils.SplitLine)
	for _, q := range qs {
		p, a := q.para, q.ans
		fmt.Printf(utils.TestLogFormat, p, fourSum(p.nums, p.target), a)
	}
	fmt.Println()
}
