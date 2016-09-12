package main

import "testing"

type event struct {
	id      int
	isEnter bool
}

func TestDFS(t *testing.T) {
	n2 := &node{id: 2}
	n1 := &node{id: 1}
	n0 := &node{
		id:  0,
		adj: []*node{n1, n2},
	}
	n2.adj = []*node{n0} // back edge

	events := []event{}

	enter := func(n *node) {
		events = append(events, event{n.id, true})
	}

	exit := func(n *node) {
		events = append(events, event{n.id, false})
	}

	dfs(n0, enter, exit)

	expectedEvents := []event{
		event{0, true},
		event{1, true},
		event{1, false},
		event{2, true},
		event{2, false},
		event{0, false},
	}

	if len(events) != len(expectedEvents) {
		t.Fatalf("expected %d events, got %v", len(expectedEvents), len(events))
	}

	for i := range expectedEvents {
		if events[i] != expectedEvents[i] {
			t.Fatalf("unexpected event at %d. expected %v, got %v", i, expectedEvents[i], events[i])
		}
	}
}
