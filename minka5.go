package main

import (
	"fmt"
)

var sub [8]int = [8]int{701, 301, 132, 57, 23, 10, 4, 1}

func swap(array []int, i, j int) {
	a := array[i]
	array[i] = array[j]
	array[j] = a
}

func insert_sort_k(array []int, k int) {
	for i := k; i < len(array); i++ {
		j := i
		for j - k >= 0 && array[j - k] > array[j] {
			swap(array, j - k, j)
			j -= k
		}
	}
}

func insert_sort(array []int) {
	for i := 0; i < 8; i++ {
		insert_sort_k(array, sub[i])
	}
}

func hIndex(citations []int) int {
	insert_sort(citations)
	clen := len(citations)
	for i := clen - 1; i >= 0; i-- {
		if citations[i] <= clen - i && citations[i] != 0 {
			return citations[i]
		}
	}
	return clen
}


func main() {
	arr := [...]int{4,4,0,0}
	fmt.Println(arr)
	fmt.Println(hIndex(arr[:]))	
}