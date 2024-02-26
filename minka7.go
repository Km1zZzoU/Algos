package main

import (
	"fmt"
)

var sub [3]int = [3]int{10, 4, 1}

func swap(array []int, i, j int) {
	a := array[i]
	array[i] = array[j]
	array[j] = a
}

func moove(array []int, i, j int) {
	for ; i < j; i++ {
		swap(array, i, j)
	}
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
	for i := 0; i < 3; i++ {
		insert_sort_k(array, sub[i])
	}
}                                      // всё для сортировки шела

func merge_sort(array []int) {
	lenarr := len(array)
	var spl1 int = lenarr - 3 * (lenarr / 4)
	var spl2 int = spl1 * 2
	mergesort(array, spl1, spl2)
}

func merge1(array []int, spl1, spl2, lenarr int) {
	p1 := 0
	p2 := spl1
	p3 := spl2
	for ; p3 < lenarr; p3++ {
		if array[p1] < array[p2] && p1 < spl1 {
			swap(array, p1, p3)
			p1++
		} else {
			swap(array, p2, p3)
			p2++
		}
	}
}

func merge2(array []int, spl1, spl2, lenarr int) {
	p1 := 0
	p2 := spl1
	p3 := spl2
	for ; p2 < p3 && p1 < spl1; p2++ {
		if (array[p1] < array[p3]) || (p2 == p3) {
			swap(array, p1, p2)
			p1++
		} else {
			swap(array, p3, p2)
			if p3 < lenarr - 1 {
				p3++
			}
		}			
	}	
}


func mergesort(array []int, spl1, spl2 int) {
	lenarr := len(array)
	if lenarr < 12 {
		insert_sort(array)
	} else {
		merge_sort(array[:spl1])
		merge_sort(array[spl1:spl2])

		merge1(array, spl1, spl2, lenarr)
		
		
		for ; spl1 > 3; {
			merge_sort(array[:spl1])
			merge2(array, spl1, spl2, lenarr)
			spl1 = spl1 / 2
			spl2 = spl2 / 2
		}

		insert_sort_k(array, 1)
	}
}

func main() {
	m := [...]int{4, 5, 2, 7, 4, 890, 23, 23, 6, 7, 23, 34, 35, 12, 34, 1, 5, 3, 43, 67, 3, 10, 18, 56, 43, 23, 12, 45, 65, 23, 54, 34, 234, 645, 23, 654, 67, 923, }
	merge_sort(m[:])
	fmt.Println(m)
}
