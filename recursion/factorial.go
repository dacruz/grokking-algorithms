package recursion

func factRec(n int) int {

	if n == 1 {
		return 1
	}

	return n * factRec(n-1)
}

func factIte(n int) int {
	fac := n

	for {
		if n == 1 {
			return fac
		}

		n--
		fac = n * fac
	}

	return fac
}
