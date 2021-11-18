package _97_integer_replacement

func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	ans := 0
	if n%2 == 0 {
		ans = integerReplacement(n / 2)
	} else {
		ans = ans + 1
		addAns := integerReplacement((n + 1) / 2)
		deleteAns := integerReplacement((n - 1) / 2)
		if addAns > deleteAns {
			ans = ans+deleteAns
		} else {
			ans = ans+addAns
		}
	}
	return 1 + ans
}
