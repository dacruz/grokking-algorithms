package sort

// O(n * log n) - using pivot as the middle position.
// The same big O applies if we use a random pivot, but it becomes O(nˆ2) if we always take the 1st position.
// Worst case: O(nˆ2)
// Avg case: O(n * log n)
// TODO - implement it using three-way partition and equal keys (not mentioned on the book)
func quick(ii []int) []int {

	if len(ii) <= 1 {
		return ii
	}

	ll := make([]int, 0)
	rr := make([]int, 0)
	p := len(ii) / 2
	for i := range ii {
		if i == p {
			continue
		}
		if ii[i] > ii[p] {
			rr = append(rr, ii[i])
		} else {
			ll = append(ll, ii[i])
		}
	}

	return append(append(quick(ll), ii[p]), quick(rr)...)

}
