package graph

type node struct {
	value interface{}
}

type vertice struct {
	n *node
	w uint
}

type graph map[*node][]vertice

type void struct{}

type visitedSet struct {
	m map[*node]void
}

func newVisitedSet() visitedSet {
	return visitedSet{
		m: map[*node]void{},
	}
}

func (v *visitedSet) add(n *node) {
	v.m[n] = void{}
}

func (v *visitedSet) contains(n *node) bool {
	_, ok := v.m[n]
	return ok
}

type parentSet struct {
	m map[*node]*node
}

func newParentSet() parentSet {
	return parentSet{
		m: map[*node]*node{},
	}
}

func (s *parentSet) set(c *node, p *node) {
	s.m[c] = p
}

func (s *parentSet) get(n *node) *node {
	return s.m[n]
}

func (ps *parentSet) backtrack(s, e *node) []*node {
	path := []*node{e}
	n := e

	for {
		p := ps.get(n)
		if p == nil {
			break
		}

		path = append(path, p)
		n = p
	}

	rev := make([]*node, len(path), len(path))
	for i, j := 0, len(path)-1; i < len(path); i, j = i+1, j-1 {
		rev[i] = path[j]
	}

	return rev
}
