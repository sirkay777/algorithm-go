package dijkstra

const UintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
const MaxInt = 1<<(UintSize-1) - 1          // 1<<31 - 1 or 1<<63 - 1

func findLowestCostNode(costs map[string]int, processed map[string]bool) string {
	var lowestCost int = MaxInt
	lowestCostNode := ""
	for k, v := range costs {
		_, ok := processed[k]
		if v < lowestCost && !ok {
			lowestCost = v
			lowestCostNode = k
		}
	}
	return lowestCostNode
}

func dijkstra() (int, []string) {
	graph := make(map[string]map[string]int)
	graph["start"] = make(map[string]int)
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2
	graph["a"] = make(map[string]int)
	graph["a"]["fin"] = 1
	graph["b"] = make(map[string]int)
	graph["b"]["a"] = 3
	graph["b"]["fin"] = 5
	graph["fin"] = make(map[string]int)

	costs := make(map[string]int)
	costs["a"] = 6
	costs["b"] = 2
	costs["fin"] = MaxInt

	parents := make(map[string]string)
	parents["a"] = "start"
	parents["b"] = "start"
	parents["fin"] = ""

	processed := make(map[string]bool)

	node := findLowestCostNode(costs, processed)
	for node != "" {
		for k, v := range graph[node] {
			newCost := costs[node] + v
			if newCost < costs[k] {
				costs[k] = newCost
				parents[k] = node
			}
		}
		processed[node] = true
		node = findLowestCostNode(costs, processed)
	}
	path := []string{"fin"}
	parent := parents["fin"]
	for parent != "start" {
		path = append([]string{parent}, path...)
		parent = parents[parent]
	}
	path = append([]string{"start"}, path...)
	return costs["fin"], path
}
