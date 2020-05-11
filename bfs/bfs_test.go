package bfs

import "testing"

func TestBFSSearch(t *testing.T) {
	graph := makeGraph()
	person, result := bfsSearch(graph, "you")
	if !result {
		t.Error("should found a seller")
	}
	if person != "thom" {
		t.Error("seller should be thom")
	}
}
