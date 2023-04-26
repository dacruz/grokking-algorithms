package binarysearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name  string
		list  []int
		item  int
		found bool
		index int
	}{
		{name: "find on the left", list: []int{1, 2, 3}, item: 1, found: true, index: 0},
		{name: "find in the middle", list: []int{1, 2, 3}, item: 2, found: true, index: 1},
		{name: "find on the right", list: []int{1, 2, 3}, item: 3, found: true, index: 2},
		{name: "search on odd sized list", list: []int{1, 2, 3, 4, 5, 6, 7}, item: 2, found: true, index: 1},
		{name: "search on even sized list", list: []int{1, 2, 3, 4, 5, 6}, item: 2, found: true, index: 1},
		{name: "search on empty list", list: []int{}, item: 3, found: false, index: 0},
		{name: "search item not on the list", list: []int{1, 2, 3}, item: 10, found: false, index: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			i, ok := search(tc.item, tc.list)
			assert.Equal(t, tc.found, ok)
			assert.Equal(t, tc.index, i)
		})
	}
}
