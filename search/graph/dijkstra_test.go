package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var book = &node{value: "Book"}
var lp = &node{value: "LP"}
var poster = &node{value: "Poster"}
var guitar = &node{value: "Guitar"}
var drums = &node{value: "Drums"}
var piano = &node{value: "Piano"}

var dg = graph{
	book: {
		vertice{n: lp, w: 5},
		vertice{n: poster, w: 0},
	},
	lp: {
		vertice{n: guitar, w: 15},
		vertice{n: drums, w: 20},
	},
	poster: {
		vertice{n: guitar, w: 30},
		vertice{n: drums, w: 35},
	},
	guitar: {
		vertice{n: piano, w: 20},
	},
	drums: {
		vertice{n: piano, w: 10},
	},
	piano: {},
}

func TestDijkstra(t *testing.T) {
	type args struct {
		start *node
		end   *node
		graph graph
	}
	tests := []struct {
		name  string
		args  args
		path  []*node
		cost  uint
		found bool
	}{
		{
			name: "find cheapest",
			args: args{
				start: book,
				end:   piano,
				graph: dg,
			},
			path:  []*node{book, lp, drums, piano},
			cost:  35,
			found: true,
		},
		{
			name: "avoid loops path exists",
			args: args{
				start: guitar,
				end:   poster,
				graph: graph{
					guitar: {vertice{n: piano, w: 1}},
					piano:  {vertice{n: guitar, w: 1}, vertice{n: poster, w: 10}},
					poster: {vertice{n: guitar, w: 1}},
				},
			},
			path:  []*node{guitar, piano, poster},
			cost:  11,
			found: true,
		},
		{
			name: "path does not exists",
			args: args{
				start: guitar,
				end:   poster,
				graph: dg,
			},
			path:  []*node{},
			found: false,
		},
		{
			name: "avoid loops path doesn't exist",
			args: args{
				start: guitar,
				end:   poster,
				graph: graph{
					guitar: {vertice{n: piano, w: 1}},
					piano:  {vertice{n: guitar, w: 1}},
					poster: {vertice{n: guitar, w: 1}},
				},
			},
			path:  []*node{},
			found: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, cost, found := dijkstra(tt.args.start, tt.args.end, tt.args.graph)
			assert.Equalf(t, tt.path, path, "dijkstra(%v, %v, %v)", tt.args.start, tt.args.end, tt.args.graph)
			assert.Equalf(t, tt.cost, cost, "dijkstra(%v, %v, %v)", tt.args.start, tt.args.end, tt.args.graph)
			assert.Equalf(t, tt.found, found, "dijkstra(%v, %v, %v)", tt.args.start, tt.args.end, tt.args.graph)
		})
	}
}
