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

func solve(input string) ([]string, []string) {

	grMap := make(map[string]*graph)
	for _, s := range strings.Split(input, "\n") {
		name := strings.Split(s, ": ")[0]
		g := &graph{
			name,
			make([]*graph, 0),
		}
		grMap[name] = g
	}
	for _, s := range strings.Split(input, "\n") {
		ind := strings.IndexAny(s, ": ")
		fns := strings.Split(s[ind+2:], ", ")
		for _, fn := range fns {
			grMap[s[:ind]].callin = append(grMap[s[:ind]].callin, grMap[fn])
		}
	}
	time := 1
	timeMap := make(map[string]int, len(grMap))
	recurs := make([]string, 0)
	for s, g := range grMap {
		if timeMap[s] == 0 {
			timeMap[s] = -1
			time = setOne(g, timeMap, time, &recurs)
		}
	}
	var highest string
	for s, i := range timeMap {
		if i == len(timeMap)+1 {
			highest = s
			break
		}
	}
	component := []string{highest}
	for s := range timeMap {
		if slices.Contains(grMap[s].callin, grMap[highest]) {
			component = append(component, s)
		}
	}
	recs := make([]string, 0)
	trace := make([]string, 0)
	for s, g := range grMap {
		if rec(g, s, &trace) {
			recs = append(recs, s)
		}
		trace = []string{}
	}
	return component, recs
}

func main() {
	input := "0: 1, 3, 4\n" +
		"1: 2, 7\n" +
		"2: 5\n" +
		"3: 3, 7\n" +
		"4: 0, 2, 6\n" +
		"5: 6\n" +
		"6: 2\n" +
		"7: 1"
	fmt.Println(solve(input))
}

func rec(g *graph, s string, trace *[]string) bool {
	for _, fn := range g.callin {
		if !slices.Contains(*trace, fn.name) {
			*trace = append(*trace, fn.name)
			if s == fn.name || rec(fn, s, trace) {
				return true
			}
		}
	}
	return false
}

func setOne(g *graph, timeMap map[string]int, time int, recurs *[]string) int {
	for _, f := range g.callin {
		if timeMap[f.name] == 0 {
			timeMap[g.name] = -1
			time = setOne(f, timeMap, time, recurs)
		}
	}
	timeMap[g.name] = time
	time++
	return time
}
