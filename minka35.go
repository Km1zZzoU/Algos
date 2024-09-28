package main

import (
  "fmt"
)

type unionFind struct {
	groups []int
	rank   []int
}

func (uf unionFind) find(n int) int {
	if uf.groups[n] == n {
		return n
	} else {
		group := uf.find(uf.groups[n])
		uf.groups[n] = group
		return group
	}
}

func (uf unionFind) union(n1, n2 int) bool {
	if uf.find(n1)+uf.find(n2) == 0 {
		uf.groups[n1] = n1
		uf.groups[n2] = n1
		uf.rank[n1] = 1
	} else if uf.find(n1)*uf.find(n2) == 0 {
		n := max(uf.find(n2), uf.find(n1))
		uf.groups[n1] = n
		uf.groups[n2] = n
	} else if uf.find(n1) == uf.find(n2) {
		return true
	} else {
		if uf.rank[uf.find(n1)] < uf.rank[uf.find(n2)] {
			uf.groups[uf.find(n1)] = uf.find(n2)
		} else {
			uf.groups[uf.find(n2)] = uf.find(n1)
			if uf.rank[uf.find(n2)] == uf.rank[uf.find(n1)] {
				uf.rank[uf.find(n2)]++
			}
		}
	}
	return false
}

func construct(n int) unionFind {
	return unionFind{make([]int, n), make([]int, n)}
}

func maxim(arr [][]int) int {
	max := 3
	for _, duo := range arr {
		for _, x := range duo {
			if max < x {
				max = x
			}
		}
	}
	return max
}

func findRedundantConnection(edges [][]int) []int {
	uf := construct(maxim(edges))
	for _, edge := range edges {
		if uf.union(edge[0], edge[1]) {
			return edge
		}
	}
	return nil
}

func main() {
	edges := [][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{1, 4}, []int{1, 5}}
	fmt.Println(findRedundantConnection(edges))
}
