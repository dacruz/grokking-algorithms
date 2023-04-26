package selectionsort

func sort(ll []int) []int {
	sl := make([]int, 0)

	for len(ll) > 0 {
		si := findSmallerIndex(ll)
		sl = append(sl, ll[si])
		ll = append(ll[:si], ll[si+1:]...)
	}

	return sl
}

func findSmallerIndex(nn []int) int {
	si := 0
	for i := range nn {
		if nn[i] < nn[si] {
			si = i
		}
	}

	return si
}
