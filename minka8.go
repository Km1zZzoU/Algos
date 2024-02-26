package main

import "fmt"

func sumSlice(slice []int) int {
    sum := 0
    for _, num := range slice {
        sum += num
    }
    return sum
}

func isIdealPermutation(nums []int) bool {
  local  := 0
  global := 0
  n      := len(nums)
  counted := make([]int, n)
  for i := 0; i < n - 1; i++ {
    counted[nums[i]]++
    local += counted[nums[i] + 1]
    global += sumSlice(counted[nums[i]:])
  }
  return global == local
}

func main() {
	m := [3]int{1, 0, 2}
	fmt.Println(isIdealPermutation(m[:]))
}