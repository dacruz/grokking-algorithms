package search

// O(log n) - short stack
func binaryIte(n int, ii []int) (int, bool) {
	l, r := 0, len(ii)-1

	for l <= r {
		m := (l + r) / 2
		if n == ii[m] {
			return m, true
		}

		if n > ii[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return 0, false
}

// O(log n) - long stack
func binaryRec(n int, ii []int) (int, bool) {
	return bRec(n, ii, 0, len(ii)-1)
}

func bRec(n int, ii []int, l int, r int) (int, bool) {
	if l > r {
		return 0, false
	}

	m := (l + r) / 2
	if ii[m] == n {
		return m, true
	}

	if n > ii[m] {
		return bRec(n, ii, m+1, r)
	} else {
		return bRec(n, ii, l, m-1)
	}

}
