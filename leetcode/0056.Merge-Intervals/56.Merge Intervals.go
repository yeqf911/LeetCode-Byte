package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	index := -1

	for _, intv := range intervals {
		if index == -1 {
			index++
			res = append(res, intv)
		} else if intv[0] > res[index][1] {
			index++
			res = append(res, intv)
		} else {
			if intv[1] > res[index][1] {
				res[index][1] = intv[1]
			}
		}
	}
	return res
}
