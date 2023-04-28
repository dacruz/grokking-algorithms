package sort

// O(nˆ2) - this function has only one loop, but the func findSmallerIndex brings another nested look over **n** into the equation
func selection(ll []int) []int {
	sl := make([]int, 0)

	for len(ll) > 0 {
		si := findSmallerIndex(ll)
		sl = append(sl, ll[si])
		ll = append(ll[:si], ll[si+1:]...)
	}

	return sl
}

// O(n)
func findSmallerIndex(nn []int) int {
	si := 0
	for i := range nn {
		if nn[i] < nn[si] {
			si = i
		}
	}

	return si
}
