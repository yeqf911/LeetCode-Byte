package leetcode

func generateParenthesis(n int) []string {
	var traceback func(int, int)
	var res []string
	var comb []byte
	traceback = func(left, right int) {
		if left < 0 || right < 0 {
			return
		}

		// 此处使用小于符号，是因为right和left都是剩余可使用的括号数量，要判定左括号比右括号使用的多，需要判定左括号剩余的比右括号少
		if right < left {
			return
		}

		if left == 0 && right == 0 {
			res = append(res, string(comb))
			return
		}

		comb = append(comb, '(')
		traceback(left-1, right)
		comb = comb[:len(comb)-1]

		comb = append(comb, ')')
		traceback(left, right-1)
		comb = comb[:len(comb)-1]
	}

	traceback(n, n)
	return res
}
