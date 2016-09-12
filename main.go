package main

type nodeState int

const (
	unvisited nodeState = iota
	visited
	processed
)

type node struct {
	id    int
	adj   []*node
	state nodeState
}

// assumes all nodes are in unvisited state
func dfs(n *node, enterAction func(*node), exitAction func(*node)) {
	q := []*node{n}
	n.state = visited

	for len(q) != 0 {
		cn := q[len(q)-1]
		if cn.state == processed {
			q = q[0 : len(q)-1]
			exitAction(cn)
			continue
		}
		enterAction(cn)

		//push neighbors
		for i := len(cn.adj) - 1; i >= 0; i-- {
			nn := cn.adj[i]
			if nn.state != unvisited {
				continue
			}
			nn.state = visited
			q = append(q, nn)
		}

		cn.state = processed
	}
}
