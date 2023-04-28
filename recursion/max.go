package recursion

func maxRec(nn []int) int {
	if len(nn) == 0 {
		return -1
	}

	m0 := nn[0]
	m1 := maxRec(nn[1:])

	if m0 > m1 {
		return m0
	} else {
		return m1
	}

}

func maxIte(nn []int) int {
	max := -1
	for _, n := range nn {
		if n > max {
			max = n
		}
	}
	return max
}
