package main

import (
	"errors"
	"fmt"
	"slices"
	"sync"
)

type g struct {
	bef   []int
	elems []int
}

type SortResult struct {
	index int
	items []int
}

func sortItems(n int, m int, group []int, beforeItems [][]int) (sortItems []int) {
	var (
		groups     = make([]*g, m+n)
		countGroup = m
	)
	for i := 0; i < n; i++ {
		if group[i] < 0 {
			group[i] = countGroup
			countGroup++
		}
	}
	for i := 0; i < n; i++ {
		if groups[group[i]] == nil {
			groups[group[i]] = &g{
				bef:   make([]int, 0),
				elems: append(make([]int, 0), i),
			}
		} else {
			groups[group[i]].elems = append(groups[group[i]].elems, i)
		}
		for j := 0; j < len(beforeItems[i]); j++ {
			el := beforeItems[i][j]
			if group[el] != group[i] && !slices.Contains(groups[group[i]].bef, group[el]) {
				groups[group[i]].bef = append(groups[group[i]].bef, group[el])
			}
			if group[i] != group[el] {
				beforeItems[i] = append(beforeItems[i][:j], beforeItems[i][j+1:]...)
				j--
			}
		}
	}
	countGroup = 0
	s2b := make(map[int]int)
	b2s := make(map[int]int)
	for i, gr := range groups {
		if gr != nil {
			b2s[i] = countGroup
			s2b[countGroup] = i
			countGroup++
		}
	}

	priorGroups := topSortGroups(&groups, countGroup, &b2s, &s2b)
	if priorGroups == nil {
		return nil
	}

	var wg sync.WaitGroup
	results := make(chan SortResult, len(priorGroups))

	for i := 0; i < len(priorGroups); i++ {
		wg.Add(1)
		go processGroup(priorGroups[i], &groups[s2b[i]].elems, &beforeItems, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	sortedResults := make(map[int][]int)

	for result := range results {
		if result.items != nil {
			sortedResults[result.index] = result.items
		}
	}

	for i := 1; i <= countGroup; i++ {
		items := sortedResults[i]
		if items == nil {
			return nil
		}
		sortItems = append(sortItems, items...)
	}

	return sortItems
}

func topSortGroups(groups *[]*g, count int, b2s, s2b *map[int]int) []int {
	var (
		numberGroup = 1
		err         error
		priority    = make([]int, count)
	)
	for i := 0; i < count; i++ {
		if priority[i] == 0 {
			priority[i] = -1
			numberGroup, err = giveNumberGroup(groups, (*s2b)[i], numberGroup, &priority, b2s)
			if err != nil {
				return nil
			}
		}
	}
	return priority
}

func giveNumberGroup(groups *[]*g, nameGroup, numberGroup int, priority *[]int, b2s *map[int]int) (int, error) {
	if len((*groups)[nameGroup].bef) == 0 && (*priority)[(*b2s)[nameGroup]] < 1 {
		(*priority)[(*b2s)[nameGroup]] = numberGroup
		numberGroup++
		return numberGroup, nil
	}

	for _, bef := range (*groups)[nameGroup].bef {
		if bef != nameGroup {
			if (*priority)[(*b2s)[bef]] == -1 {
				return -1, errors.New("cycle")
			} else if (*priority)[(*b2s)[bef]] == 0 {
				(*priority)[(*b2s)[nameGroup]] = -1
				var err error
				numberGroup, err = giveNumberGroup(groups, bef, numberGroup, priority, b2s)
				if err != nil {
					return -1, err
				}
			}
		}
	}
	(*priority)[(*b2s)[nameGroup]] = numberGroup
	numberGroup++
	return numberGroup, nil
}

func processGroup(num int, items *[]int, beforeItems *[][]int, results chan<- SortResult, wg *sync.WaitGroup) {
	defer wg.Done()
	locSortItems := make([]int, 0)

	if items != nil {
		sortNum := topSort(items, beforeItems)
		if sortNum == nil {
			results <- SortResult{
				index: -1,
				items: nil,
			}
			return
		}
		for _, el := range sortNum {
			locSortItems = append(locSortItems, el)
		}
	}
	results <- SortResult{
		index: num,
		items: locSortItems,
	}
}

func topSort(items *[]int, beforeItems *[][]int) []int {
	var (
		count        = len(*items)
		numberNumber = 1
		priority     = make([]int, count)
		err          error
	)
	for i := 0; i < count; i++ {
		if priority[i] == 0 {
			priority[i] = -1
			numberNumber, err = giveNumber(items, i, numberNumber, &priority, beforeItems)
			if err != nil {
				return nil
			}
		}
	}
	swapsort(items, &priority, 0, count-1)
	return *items
}

func giveNumber(items *[]int, inumber, numberNumber int, priority *[]int, beforeItems *[][]int) (int, error) {
	if (*beforeItems)[(*items)[inumber]] == nil && (*priority)[inumber] < 1 {
		(*priority)[inumber] = numberNumber
		numberNumber++
		return numberNumber, nil
	}

	for _, bef := range (*beforeItems)[(*items)[inumber]] {
		if bef != (*items)[inumber] {
			ibef := slices.Index(*items, bef)
			if (*priority)[ibef] == -1 {
				return -1, errors.New("cycle")
			} else if (*priority)[ibef] == 0 {
				(*priority)[ibef] = -1
				var err error
				numberNumber, err = giveNumber(items, ibef, numberNumber, priority, beforeItems)
				if err != nil {
					return -1, err
				}
			}
		}
	}
	(*priority)[inumber] = numberNumber
	numberNumber++
	return numberNumber, nil
}

func swapsort(items *[]int, priority *[]int, low, high int) {
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
			(*items)[start], (*items)[end] = (*items)[end], (*items)[start]
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
				(*items)[start], (*items)[end] = (*items)[end], (*items)[start]
				(*priority)[start], (*priority)[end] = (*priority)[end], (*priority)[start]

				start++
				end--
			}
		}
	}
	p := max(start, end)
	swapsort(items, priority, low, p)
	swapsort(items, priority, p, high)
}

func main() {
	/*
		n := 10000
		m := 5000

		//             0   1   2  3  4  5  6  7
		group := make([]int, n)
		for i := 0; i < n; i++ {
			group[i] = i / 2
		}
		beforeItems := make([][]int, n)
		for i := 0; i < n; i++ {
			beforeItems[i] = make([]int, 1)
			for j := 0; j < 1; j++ {
				beforeItems[i][j] = 0
			}
		}*/
	n := 5
	m := 5
	//              0  1  2  3  4  5  6  7
	group := []int{2, 0, -1, 3, 0}
	beforeItems := [][]int{
		{2, 1, 3}, // 0
		{2, 4},    // 1
		{},        // 2
		{},        // 3
		{},
	}
	fmt.Println(sortItems(n, m, group, beforeItems))
}
