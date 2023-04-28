package recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var maxTests = []struct {
	name   string
	list   []int
	expect int
}{
	{name: "stops at list with one element", list: []int{1}, expect: 1},
	{name: "fails on empty list", list: []int{}, expect: -1},
	{name: "find max", list: []int{1, 2, 10, 3, 4}, expect: 10},
}

func TestRecursiveMax(t *testing.T) {
	execMax(t, maxRec)
}

func TestIterativeMax(t *testing.T) {
	execMax(t, maxIte)
}

func execMax(t *testing.T, f func([]int) int) {
	for _, tc := range maxTests {
		t.Run(tc.name, func(t *testing.T) {
			f := f(tc.list)
			assert.Equal(t, tc.expect, f)
		})
	}
}
