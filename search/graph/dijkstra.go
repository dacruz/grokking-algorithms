package graph

import "math"

const INF = math.MaxUint

// O(Vˆ2) - this Big O is complex. See more: https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
func dijkstra(s *node, e *node, g graph) ([]*node, uint, bool) {

	vs := newVisitedSet()
	ps := newParentSet()

	costs := make(map[*node]uint)
	costs[s] = 0

	n := s
	var prev *node
	for n != nil {
		for _, v := range g[n] {
			cost, ok := costs[v.n]
			if !ok {
				cost = INF
			}

			if cost > costs[n]+v.w {
				ps.set(v.n, n)
				costs[v.n] = costs[n] + v.w
			}
		}

		vs.add(n)
		prev = n
		n = cheaper(costs, vs)
	}

	if prev != e {
		return []*node{}, 0, false
	}

	return ps.backtrack(s, e), costs[e], true
}

func cheaper(costs map[*node]uint, vs visitedSet) *node {

	var cn *node
	var lowest uint

	lowest = INF
	for n, c := range costs {
		if lowest > c && !vs.contains(n) {
			cn = n
			lowest = c
		}
	}

	return cn
}
