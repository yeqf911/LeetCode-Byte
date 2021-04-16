package leetcode

import (
	"LeetCode-Byte/structure"
	"reflect"
	"testing"
)

func Test_addTwoNumbers1(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "0001",
			args: args{
				structure.Array2Linklist([]int{2, 4, 3}),
				structure.Array2Linklist([]int{5, 6, 4}),
			},
			want: structure.Array2Linklist([]int{7, 0, 8}),
		},
		{
			name: "0002",
			args: args{
				structure.Array2Linklist([]int{9, 9, 9, 9, 9, 9, 9}),
				structure.Array2Linklist([]int{9, 9, 9, 9}),
			},
			want: structure.Array2Linklist([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
