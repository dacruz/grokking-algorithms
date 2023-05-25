package graph

func breadthFirst(s *node, e *node, g graph) ([]*node, bool) {
	q := []*node{s}

	v := newVisitedSet()
	p := newParentSet()

	for len(q) > 0 {
		n := q[0]
		if v.contains(n) {
			q = q[1:]
			continue
		}
		v.add(n)

		if n == e {
			return p.backtrack(s, e), true
		} else {
			q = q[1:]
			for _, ver := range g[n] {
				if !v.contains(ver.n) {
					p.set(ver.n, n)
					q = append(q, ver.n)
				}
			}
		}
	}

	return []*node{}, false
}
