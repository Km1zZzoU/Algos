package main

import "fmt"

type unionFind struct {
	groups []int
	rank   []int
}

type task struct {
	name     string
	deadLine int
	price    int
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

func (uf unionFind) union(x, y int) {
	a := uf.find(x)
	b := uf.find(y)
	if uf.rank[a] > uf.rank[b] {
		uf.groups[b] = a
	} else {
		uf.groups[a] = b
		if uf.rank[a] == uf.rank[b] {
			uf.rank[b]++
		}
	}
}

func construct(n int) unionFind {
	groups := make([]int, n)
	for i := range groups {
		groups[i] = i
	}
	return unionFind{groups, make([]int, n)}
}

func solve(tasks []task) (sortasks []string) {
	len := len(tasks)
	rightPointer := len - 1
	sortasks = make([]string, len)
	uf := construct(len + 1)
	mapa := make([]int, len+1)

	for i := 0; i <= len; i++ {
		mapa[i] = i
	}

	for _, t := range tasks {
		place := mapa[uf.find(t.deadLine)]
		if place > 0 {
			sortasks[place-1] = t.name
			mapa[uf.find(place)] = uf.find(place - 1)
			mapa[uf.find(place-1)] = uf.find(place - 1)
			uf.union(place, place-1)
			// uf.groups[place] = uf.find(place - 1)
			if place-1 == rightPointer {
				rightPointer--
			}
		} else {
			sortasks[rightPointer] = t.name
			mapa[uf.find(rightPointer+1)] = uf.find(rightPointer)
			mapa[uf.find(rightPointer)] = uf.find(rightPointer)
			uf.union(rightPointer+1, rightPointer)
			rightPointer--
		}
	}
	return sortasks
}

func main() {
	tasks := []task{
		{"D", 3, 50},
		{"C", 1, 30},
		{"A", 3, 25},
		{"E", 3, 20},
		{"B", 4, 10},
	}
	fmt.Println(solve(tasks))

	tasks = []task{
		{"G", 5, 40},
		{"H", 1, 35},
		{"J", 2, 30},
		{"I", 4, 25},
		{"F", 2, 15},
	}
	fmt.Println(solve(tasks))

	tasks = []task{
		{"D", 2, 50},
		{"C", 2, 40},
		{"A", 2, 35},
		{"E", 4, 10},
	}
	fmt.Println(solve(tasks))

	tasks = []task{
		{"D", 5, 50},
		{"D2", 5, 45},
		{"C", 5, 35},
		{"A", 5, 25},
		{"E", 2, 20},
		{"B", 1, 10},
	}
	fmt.Println(solve(tasks))
}
