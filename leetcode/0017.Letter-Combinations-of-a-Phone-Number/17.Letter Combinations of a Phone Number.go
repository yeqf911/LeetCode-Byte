package leetcode

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	keyboard := map[byte][]byte{
		'2': []byte("abc"),
		'3': []byte("def"),
		'4': []byte("ghi"),
		'5': []byte("jkl"),
		'6': []byte("mno"),
		'7': []byte("pqrs"),
		'8': []byte("tuv"),
		'9': []byte("wxyz"),
	}

	var res []string
	var traceback func([]byte, int)
	traceback = func(sequence []byte, idx int) {
		if len(sequence) == len(digits) {
			res = append(res, string(sequence))
			return
		}

		d := digits[idx]
		bs, _ := keyboard[d]
		for _, c := range bs {
			sequence = append(sequence, c)
			traceback(sequence, idx+1)
			sequence = sequence[:len(sequence)-1]
		}
	}

	sq := make([]byte, 0)
	traceback(sq, 0)
	return res
}
