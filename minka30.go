package main

import (
	"errors"
	"fmt"
	"slices"
)

type elem struct {
	item  int
	group int
	bef   []int
}

type graph struct {
	num int
	bef []*graph
}

func sortItems(n int, m int, group []int, beforeItems [][]int) (sortItems []int) {
	var (
		countGroup = m
		elems      = make([]elem, n)
		elems2     = make([][]*graph, n+m)
		elemsline  = make([]*graph, 0)
	)
	for i := 0; i < n; i++ {
		if group[i] < 0 {
			group[i] = countGroup
			countGroup++
		}
		elems[i] = elem{
			item:  i,
			group: group[i],
			bef:   beforeItems[i],
		}
		gr := graph{
			num: i,
			bef: nil,
		}
		elemsline = append(elemsline, &gr)
		if elems2[group[i]] == nil {
			arr := make([]*graph, 0)
			elems2[group[i]] = append(arr, &gr)
		} else {
			elems2[group[i]] = append(elems2[group[i]], &gr)
		}
	}
	groups := make([]*graph, countGroup)
	for i := range groups {
		groups[i] = &graph{
			num: i,
			bef: nil,
		}
	}
	for _, e := range elems {
		for _, b := range e.bef {
			if groups[e.group] != groups[elems[b].group] {
				if groups[e.group].bef == nil {
					groups[e.group].bef = []*graph{groups[elems[b].group]}
				} else if !slices.Contains(groups[e.group].bef, groups[elems[b].group]) {
					groups[e.group].bef = append(groups[e.group].bef, groups[elems[b].group])
				}
			}
		}
	}
	sortGroups := topsort(groups, n+m)
	if sortGroups == nil {
		return nil
	}

	for _, grp := range sortGroups {
		//для каждой группы
		for _, grf := range elems2[grp.num] {
			// для каждого графа из группы
			for _, numbef := range elems[grf.num].bef {
				if grp.num == elems[grf.num].group {
					grf.bef = append(grf.bef, elemsline[numbef])
				}
			}
		}
		if elems2[grp.num] != nil {
			sortNum := topsort(elems2[grp.num], n+m)
			if sortNum == nil {
				return nil
			}
			for _, num := range sortNum {
				sortItems = append(sortItems, num.num)
			}
		}
	}
	return sortItems
}

func topsort(Nodes []*graph, nm int) (sortNodes []*graph) {
	var (
		count      = len(Nodes)
		numberNode = 1
		priotity   = make([]int, nm)
		compprio   = make([]int, 0)
		err        error
	)
	for i := 0; i < count; i++ {
		if priotity[Nodes[i].num] == 0 {
			priotity[Nodes[i].num] = -1
			numberNode, err = giveNumber(Nodes[i], numberNode, &priotity)
			if err != nil {
				return nil
			}
		}
	}
	for _, p := range priotity {
		if p > 0 {
			compprio = append(compprio, p)
		}
	}
	swapsort(&Nodes, &compprio, 0, count-1)
	return Nodes
}

func swapsort(Nodes *[]*graph, priority *[]int, low, high int) {
	var (
		pivot = (*priority)[(high+low)/2]
		start = low
		end   = high
		l     = high - low + 1
	)
	if l < 2 {
		return
	} else if l == 2 {
		if (*priority)[low] > (*priority)[high] {
			(*Nodes)[start], (*Nodes)[end] = (*Nodes)[end], (*Nodes)[start]
		}
		return
	}
	for start < end {
		var (
			a = (*priority)[start]
			b = (*priority)[end]
		)
		if a < pivot {
			if b < pivot {
				start++
			} else {
				start++
				end--
			}
		} else {
			if pivot < b {
				end--
			} else {
				(*Nodes)[start], (*Nodes)[end] = (*Nodes)[end], (*Nodes)[start]
				(*priority)[start], (*priority)[end] = (*priority)[end], (*priority)[start]

				start++
				end--
			}
		}
	}
	p := max(start, end)
	swapsort(Nodes, priority, low, p)
	swapsort(Nodes, priority, p, high)
}

func giveNumber(Node *graph, numberNode int, priority *[]int) (int, error) {
	if Node.bef == nil && (*priority)[Node.num] < 1 {
		(*priority)[Node.num] = numberNode
		numberNode++
		return numberNode, nil
	}

	for _, bef := range Node.bef {
		if bef != Node {
			if (*priority)[bef.num] == -1 {
				return -1, errors.New("cycle")
			} else if (*priority)[bef.num] == 0 {
				(*priority)[Node.num] = -1
				var err error
				numberNode, err = giveNumber(bef, numberNode, priority)
				if err != nil {
					return -1, err
				}
			}
		}
	}
	(*priority)[Node.num] = numberNode
	numberNode++
	return numberNode, nil
}

func main() {
	n := 5
	m := 3

	//             0   1   2  3  4  5  6  7
	group := []int{0, 0, 2, 1, 0}
	beforeItems := [][]int{
		{3},       // 0
		{},        // 1
		{},        // 2
		{},        // 3
		{1, 3, 2}, // 4
	}
	fmt.Println(sortItems(n, m, group, beforeItems))
}
