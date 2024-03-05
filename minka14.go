package main

import (
	"fmt"
	"math/rand"
)

func swap(array []int, i, j int) {
	a := array[i]
	array[i] = array[j]
	array[j] = a
}

func sortArray(nums []int, split int) []int {
	qsort(nums, split)
	return nums
}

func qsort(nums []int, split int) {
	if len(nums) > 1 {
		j := len(nums) - 1
		ind := rand.Intn(j) + 1
		pivot := nums[ind]
		swap(nums, ind, 0)
		i := 1
		for i <= j {
			if nums[i] < pivot {
				if nums[j] < pivot {
					i++
				} else {
					i++
					j--
				}
			} else {
				if nums[j] > pivot {
					j--
				} else {
					swap(nums, i, j)
					i++
					j--
				}
			}
		}
		swap(nums, j, 0)
		if j > split {
			sortArray(nums[:j], split)
		} else if j < split {
			sortArray(nums[j:], split-j)
		}

	}
}

func neft(m []int) int {
	N := len(m)
	var split int = N / 2
	m = sortArray(m, split)
	return m[split]
}

func main() {
	array := []int{7, 4, 5, 2, 6, 8, 3, 9, 10, 2, 2, 10, 10, 10}
	fmt.Println(neft(array))
}
