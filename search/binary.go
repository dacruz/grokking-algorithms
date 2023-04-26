package binarysearch

func search(n int, ii []int) (int, bool) {
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
