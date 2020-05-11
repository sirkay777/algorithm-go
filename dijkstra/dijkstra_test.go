package dijkstra

import "testing"

func TestDijkstra(t *testing.T) {
	expectedCost := 6
	expectedPath := []string{"start", "b", "a", "fin"}
	cost, path := dijkstra()
	if cost != expectedCost {
		t.Errorf("cost expected: %d, got: %d\n", expectedCost, cost)
	}
	if !arrSameOrder(path, expectedPath) {
		t.Errorf("path expected: %v, got: %v\n", expectedPath, path)
	}
}

func arrSameOrder(arr1 []string, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
