package selectionsort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name     string
		list     []int
		expected []int
	}{
		{name: "sort list", list: []int{2, 1, 3}, expected: []int{1, 2, 3}},
		{name: "sort already sorted list", list: []int{1, 2, 3}, expected: []int{1, 2, 3}},
		{name: "sort empty list", list: []int{}, expected: []int{}},
		{name: "sort list with duplicates", list: []int{1, 2, 1, 2, 1}, expected: []int{1, 1, 1, 2, 2}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sl := sort(tc.list)
			assert.EqualValues(t, tc.expected, sl)
		})
	}
}
