package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var you = &node{value: "you"}
var bob = &node{value: "bob"}
var alice = &node{value: "alice"}
var claire = &node{value: "claire"}
var anuj = &node{value: "anuj"}
var peggy = &node{value: "peggy"}
var thom = &node{value: "thom"}
var jonny = &node{value: "jonny"}
var mary = &node{value: "mary"}

var bfg = graph{
	you: {
		vertice{n: bob},
		vertice{n: alice},
		vertice{n: claire},
	},
	bob: {
		vertice{n: anuj},
		vertice{n: peggy},
	},
	alice: {
		vertice{n: peggy},
	},
	claire: {
		vertice{n: thom},
		vertice{n: jonny},
	},
	anuj:  {},
	peggy: {},
	thom:  {},
	jonny: {},
	mary:  {},
}

func TestBreadthFirst(t *testing.T) {
	type args struct {
		start *node
		end   *node
		graph graph
	}
	tests := []struct {
		name  string
		args  args
		found bool
		path  []*node
	}{
		{
			name: "path exists",
			args: args{
				start: you,
				end:   peggy,
				graph: bfg,
			},
			found: true,
			path:  []*node{you, alice, peggy},
		},
		{
			name: "path does not exists",
			args: args{
				start: you,
				end:   mary,
				graph: bfg,
			},
			found: false,
			path:  []*node{},
		},
		{
			name: "avoid loops path doesn't exist",
			args: args{
				start: you,
				end:   jonny,
				graph: graph{
					you:   {vertice{n: alice}},
					alice: {vertice{n: you}, vertice{n: mary}},
					jonny: {},
					mary:  {},
				},
			},
			found: false,
			path:  []*node{},
		},
		{
			name: "avoid loops path exists",
			args: args{
				start: you,
				end:   jonny,
				graph: graph{
					you:   {vertice{n: alice}},
					alice: {vertice{n: you}, vertice{n: mary}},
					mary:  {vertice{n: jonny}},
					jonny: {},
				},
			},
			found: true,
			path:  []*node{you, alice, mary, jonny},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, exists := breadthFirst(tt.args.start, tt.args.end, tt.args.graph)
			assert.Equalf(t, tt.found, exists, "breadthFirst(%v, %v, %v)", tt.args.start, tt.args.end, tt.args.graph)
			assert.Equalf(t, tt.path, path, "breadthFirst(%v, %v, %v)", tt.args.start, tt.args.end, tt.args.graph)
		})
	}
}
