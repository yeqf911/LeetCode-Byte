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

func Test_removeDuplicates(t *testing.T) {
	qs := []question{
		{
			para{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			ans{want: 5},
		},
	}

	fmt.Println(utils.SplitLine)
	for _, q := range qs {
		p, a := q.para, q.ans
		fmt.Printf(utils.TestLogFormat, p, removeDuplicates(p.nums), a)
	}
	fmt.Println()
}
