package bfs

import "fmt"

type Graph map[string][]string

func makeGraph() Graph {
	graph := make(map[string][]string)
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}
	return graph
}

func personIsASeller(person string) bool {
	return person == "thom"
}

func bfsSearch(graph Graph, start string) (string, bool) {
	var queue []string
	visited := make(map[string]bool)
	queue = append(queue, start)
	for len(queue) != 0 {
		currentNode := queue[0]
		queue = queue[1:]
		_, ok := visited[currentNode]
		if !ok {
			visited[currentNode] = true
			if personIsASeller(currentNode) {
				fmt.Printf("%s is a seller\n", currentNode)
				return currentNode, true
			} else {
				queue = append(queue, graph[currentNode]...)
			}
		}
	}
	return "", false
}
