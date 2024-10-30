package main

func numTrees(n int) int {
	if n < 2 {
		return 1
	}
	arr := make([]int, n+1)
	arr[0] = 1
	arr[1] = 1
	for i := 2; i <= n; i++ {
		arr[i] = 0
		for j := 0; j < n; j++ {
			arr[i] += arr[j] * arr[i-j-1]
		}
	}
	return arr[n]
}

func main() {
	print(numTrees(3))
}
