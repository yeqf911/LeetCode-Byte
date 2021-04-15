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
	want [][]int
}

type question struct {
	para
	ans
}

func Test_threeSums(t *testing.T) {
	var qs = []question{
		{
			para{[]int{1, 2, 3, -3}},
			ans{[][]int{{1, 2, -3}}},
		},
		{
			para{[]int{-1, 0, 1, 1, 2, -1, -4, -4}},
			ans{[][]int{{-1, -1, 2}, {-1, 0, 1}}},
		},
	}
	fmt.Println(utils.SplitLine)
	for _, q := range qs {
		p, a := q.para, q.ans
		fmt.Printf(utils.TestLogFormat, p, threeSum(p.nums), a)
	}
	fmt.Println()
}
