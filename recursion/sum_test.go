package recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var sumTests = []struct {
	name   string
	list   []int
	expect int
}{
	{name: "stops at empty list", list: []int{}, expect: 0},
	{name: "calc sum", list: []int{1, 2, 3, 4}, expect: 10},
}

func TestRecursiveSum(t *testing.T) {
	execSum(t, sumRec)
}

func TestIterativeSum(t *testing.T) {
	execSum(t, sumIte)
}

func execSum(t *testing.T, f func([]int) int) {
	for _, tc := range sumTests {
		t.Run(tc.name, func(t *testing.T) {
			f := f(tc.list)
			assert.Equal(t, tc.expect, f)
		})
	}
}
