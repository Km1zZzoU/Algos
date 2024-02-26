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

func wiggleSort(nums []int)  {
	insert_sort(nums)
	arr := make([]int, len(nums))
	copy(arr, nums)
	alen := len(arr)
	middle := alen / 2
	if middle * 2 != alen {
		nums[alen - 1] = arr[0]
	}
	for i := 0; i < middle; i++ {
		nums[2 * i] = arr[middle - i - 1]
		nums[2 * i + 1] = arr[alen - i - 1]
	}
}

func main() {
	a := []int{1,5,1,1,6,4}
	wiggleSort(a)
	fmt.Println(a)
}