package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		bs := []byte(s)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		key := string(bs)
		if _, ok := m[key]; ok {
			m[key] = append(m[key], s)
		} else {
			m[key] = append([]string{s})
		}
	}

	var strArrays [][]string
	for _, v := range m {
		strArrays = append(strArrays, v)
	}
	return strArrays
}

func groupAnagrams2(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, s := range strs {
		bs := [26]int{}
		for _, r := range s {
			index := r - rune('a')
			bs[index]++
		}
		if _, ok := m[bs]; ok {
			m[bs] = append(m[bs], s)
		} else {
			m[bs] = append([]string{}, s)
		}
	}

	var res [][]string
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
