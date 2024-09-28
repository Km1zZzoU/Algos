package main

import (
	"fmt"
	"slices"
	"strings"
)

type graph struct {
	name   string
	callin []*graph
}

func solve(input string) (strongComp []string, recs []string) {
	grMap := make(map[string]*graph)
	revMap := make(map[string]*graph)
	countFunc := 0
	for _, s := range strings.Split(input, "\n") {
		countFunc++
		name := strings.Split(s, ": ")[0]
		g := &graph{
			name,
			make([]*graph, 0),
		}
		grMap[name] = g
		f := &graph{
			name,
			make([]*graph, 0),
		}
		revMap[name] = f
	}
	for _, s := range strings.Split(input, "\n") {
		ind := strings.IndexAny(s, ": ")
		fns := strings.Split(s[ind+2:], ", ")
		for _, fn := range fns {
			grMap[s[:ind]].callin = append(grMap[s[:ind]].callin, grMap[fn])
			if !slices.Contains(revMap[fn].callin, revMap[s[:ind]]) {
				revMap[fn].callin = append(revMap[fn].callin, revMap[s[:ind]])
			}
		}
	}
	time := 1
	timeMap := make(map[string]int, len(revMap))
	for s, g := range revMap {
		if timeMap[s] == 0 {
			timeMap[s] = -1
			setOne(g, &timeMap, &time)
		}
	}
	var next string
	dfsMap := make(map[string]bool, len(grMap))
	for name, t := range timeMap {
		if t == countFunc {
			next = name
		}
	}
	lowNumInCom := countFunc
	comps := make([][]string, 0)
	var newComp []string
	var lastComps []string
	dfsMap = make(map[string]bool, len(grMap))
	for lowNumInCom > 1 {
		detectRec(grMap[next], &dfsMap, &timeMap, &lowNumInCom)
		newComp = make([]string, 0)
		if lastComps == nil {
			for name, pred := range dfsMap {
				if pred {
					lastComps = append(lastComps, name)
				}
			}
			newComp = lastComps
		} else {
			for name, b := range dfsMap {
				if b && !slices.Contains(lastComps, name) {
					newComp = append(newComp, name)
					lastComps = append(lastComps, name)
				}
			}
		}
		if len(newComp) > len(strongComp) {
			strongComp = newComp
		}
		if len(newComp) > 1 || slices.Contains(grMap[newComp[0]].callin, grMap[newComp[0]]) {
			comps = append(comps, newComp)
		}
		for name, t := range timeMap {
			if t == lowNumInCom-1 {
				next = name
				break
			}
		}
	}
	for _, comp := range comps {
		for _, name := range comp {
			recs = append(recs, name)
		}
	}
	return strongComp, recs
}

func main() {
	input := "0: 4\n" +
		"1: 0, 7\n" +
		"2: 1, 4, 6\n" +
		"3: 0, 3\n" +
		"4: 0\n" +
		"5: 2\n" +
		"6: 4, 5\n" +
		"7: 1, 3"
	fmt.Println(solve(input))
}

func setOne(g *graph, timeMap *map[string]int, time *int) {
	for _, f := range g.callin {
		if (*timeMap)[f.name] == 0 {
			(*timeMap)[f.name] = -1
			setOne(f, timeMap, time)
		}
	}
	(*timeMap)[g.name] = *time
	*time++
}

func detectRec(g *graph, dfsMap *map[string]bool, timeMap *map[string]int, lowly *int) {
	if (*timeMap)[g.name] < *lowly {
		*lowly = (*timeMap)[g.name]
	}
	(*dfsMap)[g.name] = true
	for _, f := range g.callin {
		if !(*dfsMap)[f.name] {
			detectRec(f, dfsMap, timeMap, lowly)
		}
	}
}
