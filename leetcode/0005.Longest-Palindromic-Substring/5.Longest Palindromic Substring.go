package leetcode

func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i+1)
		if len(s1) > len(s2) {
			if len(res) < len(s1) {
				res = s1
			}
		} else {
			if len(res) < len(s2) {
				res = s2
			}
		}
	}
	return res
}

func palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	// 这里的r是开区间
	return s[l+1 : r]
}