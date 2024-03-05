package main

func swap(array []int, i, j int) {
	a := array[i]
	array[i] = array[j]
	array[j] = a
}

func sortColors(nums []int) {
	p1 := 0
	p2 := 0
	p3 := len(nums) - 1
	for p2 <= p3 {
		if nums[p2] == 0 {
			swap(nums, p1, p2)
			p1++
			p2++
		} else if nums[p2] == 2 {
			swap(nums, p2, p3)
			p3--
		} else {
			p2++
		}
	}
}

func main() {
	nums := []int{2, 0, 1}
	sortColors(nums)
}
