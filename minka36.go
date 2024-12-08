package main

import "fmt"

type unionFind struct {
	groups []int
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

func construct(n int) unionFind {
	groups := make([]int, n)
	for i := range groups {
		groups[i] = i
	}
	return unionFind{groups}
}

func solve(tasks []task) (sortasks []string) {
	len := len(tasks)
	rightPointer := len - 1
	sortasks = make([]string, len)
	uf := construct(len + 1)
	for _, t := range tasks {
		place := uf.find(t.deadLine)
		if place > 0 {
			sortasks[place-1] = t.name
			uf.groups[place] = uf.find(place - 1)
			if place-1 == rightPointer {
				rightPointer--
			}
		} else {
			sortasks[rightPointer] = t.name
			uf.groups[rightPointer+1] = uf.find(rightPointer)
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
