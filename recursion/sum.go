package recursion

// O(n) - short stack
func sumIte(nn []int) int {

	s := 0
	for _, n := range nn {
		s = s + n
	}

	return s
}

// O(n) - long stack
func sumRec(nn []int) int {

	if len(nn) == 0 {
		return 0
	}

	return nn[0] + sumRec(nn[1:])
}
