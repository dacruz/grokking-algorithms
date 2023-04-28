package recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var factTests = []struct {
	name   string
	num    int
	expect int
}{
	{name: "stops at 1", num: 1, expect: 1},
	{name: "fact 2", num: 2, expect: 2},
	{name: "fact 3", num: 3, expect: 6},
}

func TestRecursiveFactorial(t *testing.T) {
	execFact(t, factRec)
}

func TestIterativeFactorial(t *testing.T) {
	execFact(t, factIte)
}

func execFact(t *testing.T, f func(int) int) {
	for _, tc := range factTests {
		t.Run(tc.name, func(t *testing.T) {
			f := f(tc.num)
			assert.Equal(t, tc.expect, f)
		})
	}
}
