package main

func isBipartite(graph [][]int) bool {
	False := false
	arr := make([]*bool, len(graph))
	for i := range graph {
		if arr[i] == nil {
			arr[i] = &False
			if inob(graph, arr, i, false) {
				return false
			}
		}
	}
	return true
}

func inob(graph [][]int, arr []*bool, index int, side bool) bool {
	True := true
	False := false
	for _, node := range graph[index] {
		if arr[node] == nil {
			if side {
				arr[node] = &False
			} else {
				arr[node] = &True
			}
			if inob(graph, arr, node, !side) {
				return true
			}
		} else if *arr[node] == *arr[index] {
			return true
		}
	}
	return false
}

func main() {
	g := [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}
	print(isBipartite(g))
}
