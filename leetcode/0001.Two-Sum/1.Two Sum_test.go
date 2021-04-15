package leetcode

import (
	"LeetCode-Byte/utils"
	"fmt"
	"testing"
)

type question struct {
	para
	ans
}

type para struct {
	nums   []int
	target int
}

type ans struct {
	want []int
}

func Test_0001(t *testing.T) {
	qs := []question{
		{
			para{[]int{1, 2, 3}, 4},
			ans{[]int{0, 2}},
		},
		{
			para{[]int{2, 4, 6}, 10},
			ans{[]int{1, 2}},
		},
	}

	fmt.Printf("------------------------ Leetcode Problem ------------------------\n")

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf(utils.TestLogFormat, p, twoSum(p.nums, p.target), a)
	}

	fmt.Printf("\n")
}
