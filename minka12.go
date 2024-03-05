package main

import (
	"math/rand"
)

func swap(array []int, i, j int) {
	a := array[i]
	array[i] = array[j]
	array[j] = a
}

func sortArray(nums []int) []int {
	qsort(nums)
	return nums
}

func qsort(nums []int) {
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
		sortArray(nums[:j])
		sortArray(nums[j+1:])
	}
}

func sortArray(nums []int) []int {
	qsort(nums)
	return nums
}

func qsort(nums []int) {
	if len(nums) > 1 {
		ind := rand.Intn(len(nums)-1) + 1
		ind = 1
		pivot := nums[ind]
		swap(nums, ind, 0)
		l := 0
		h := 0
		c := 1
		for ; c < len(nums); c++ {
			if nums[c] < pivot {
				tmp := nums[c]
				nums[c] = nums[h+1]
				nums[h+1] = nums[l]
				nums[l] = tmp
				l++
				h++
			} else if nums[c] == pivot {
				swap(nums, h+1, c)
				h++
			}
		}

		qsort(nums[:l])
		qsort(nums[h+1:])
	}
}

func main() {
	array := []int{4, 5, 2, 5, 5, 8, 3, 9}
	array = sortArray(array)
}
