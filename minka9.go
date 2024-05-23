package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func grm(rows, cols int) [][]int {
	matrix := makematrix(rows, cols)
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = rand.Intn(100000)
		}
	}
	return matrix
}

func makematrix(rows, cols int) [][]int {
	result := make([][]int, rows)
	for i := range result {
		result[i] = make([]int, cols)
	}
	return result
}

func addzero(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])
	size := max(rows, cols)
	resize := 1
	for resize < size {
		resize *= 2
	}

	result := makematrix(resize, resize)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[i][j] = matrix[i][j]
		}
	}
	return result
}

func classic(matrix1, matrix2 [][]int) [][]int {
	rows := len(matrix1)
	cols := len(matrix2[0])
	result := makematrix(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for k := 0; k < len(matrix2); k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}
	return result
}

func summatrix(matrix1, matrix2 [][]int) [][]int {
	result := makematrix(len(matrix1), len(matrix1[0]))
	for i := range result {
		for j := range matrix1[i] {
			result[i][j] = matrix1[i][j] + matrix2[i][j]
		}
	}

	return result
}

func diffmatrix(matrix1, matrix2 [][]int) [][]int {
	result := makematrix(len(matrix1), len(matrix1[0]))
	for i := range result {
		for j := range matrix1[i] {
			result[i][j] = matrix1[i][j] - matrix2[i][j]
		}
	}

	return result
}

func recur(matrix1, matrix2 [][]int) [][]int {
	matrix1 = addzero(matrix1)
	matrix2 = addzero(matrix2)
	n := len(matrix1) / 2
	if n < 100 {
		return classic(matrix1, matrix2)
	}
	A := makematrix(n, n)
	B := makematrix(n, n)
	C := makematrix(n, n)
	D := makematrix(n, n)
	for i := range matrix1 {
		for j := range matrix1[i] {
			if i < n {
				if j < n {
					A[i][j] = matrix1[i][j]
				} else {
					B[i][j-n] = matrix1[i][j]
				}
			} else {
				if j < n {
					C[i-n][j] = matrix1[i][j]
				} else {
					D[i-n][j-n] = matrix1[i][j]
				}
			}
		}
	}
	E := makematrix(n, n)
	F := makematrix(n, n)
	G := makematrix(n, n)
	H := makematrix(n, n)
	for i := range matrix2 {
		for j := range matrix2[i] {
			if i < n {
				if j < n {
					E[i][j] = matrix2[i][j]
				} else {
					F[i][j-n] = matrix2[i][j]
				}
			} else {
				if j < n {
					G[i-n][j] = matrix2[i][j]
				} else {
					H[i-n][j-n] = matrix2[i][j]
				}
			}
		}
	}
	res1 := summatrix(recur(A, E), recur(B, G))
	res2 := summatrix(recur(A, F), recur(B, H))
	res3 := summatrix(recur(C, E), recur(D, G))
	res4 := summatrix(recur(C, F), recur(D, H))
	res := makematrix(2*n, 2*n)
	for i := range res {
		for j := range res[i] {
			if i < n {
				if j < n {
					res[i][j] = res1[i][j]
				} else {
					res[i][j] = res2[i][j-n]
				}
			} else {
				if j < n {
					res[i][j] = res3[i-n][j]
				} else {
					res[i][j] = res4[i-n][j-n]
				}
			}
		}
	}
	return res
}

func fastrec(matrix1, matrix2 [][]int) [][]int {
	matrix1 = addzero(matrix1)
	matrix2 = addzero(matrix2)
	n := len(matrix1) / 2
	if n < 100 {
		return classic(matrix1, matrix2)
	}
	A := makematrix(n, n)
	B := makematrix(n, n)
	C := makematrix(n, n)
	D := makematrix(n, n)
	for i := range matrix1 {
		for j := range matrix1[i] {
			if i < n {
				if j < n {
					A[i][j] = matrix1[i][j]
				} else {
					B[i][j-n] = matrix1[i][j]
				}
			} else {
				if j < n {
					C[i-n][j] = matrix1[i][j]
				} else {
					D[i-n][j-n] = matrix1[i][j]
				}
			}
		}
	}
	E := makematrix(n, n)
	F := makematrix(n, n)
	G := makematrix(n, n)
	H := makematrix(n, n)
	for i := range matrix2 {
		for j := range matrix2[i] {
			if i < n {
				if j < n {
					E[i][j] = matrix2[i][j]
				} else {
					F[i][j-n] = matrix2[i][j]
				}
			} else {
				if j < n {
					G[i-n][j] = matrix2[i][j]
				} else {
					H[i-n][j-n] = matrix2[i][j]
				}
			}
		}
	}
	P1 := fastrec(A, diffmatrix(F, H))
	P2 := fastrec(summatrix(A, B), H)
	P3 := fastrec(summatrix(C, D), E)
	P4 := fastrec(D, diffmatrix(G, E))
	P5 := fastrec(summatrix(A, D), summatrix(E, H))
	P6 := fastrec(diffmatrix(B, D), summatrix(G, H))
	P7 := fastrec(diffmatrix(A, C), summatrix(E, F))

	res1 := diffmatrix(summatrix(P6, summatrix(P5, P4)), P2)
	res2 := summatrix(P1, P2)
	res3 := summatrix(P3, P4)
	res4 := diffmatrix(summatrix(P1, P5), summatrix(P3, P7))
	res := makematrix(2*n, 2*n)
	for i := range res {
		for j := range res[i] {
			if i < n {
				if j < n {
					res[i][j] = res1[i][j]
				} else {
					res[i][j] = res2[i][j-n]
				}
			} else {
				if j < n {
					res[i][j] = res3[i-n][j]
				} else {
					res[i][j] = res4[i-n][j-n]
				}
			}
		}
	}
	return res
}

func main() {
	benchmarks := []int{}
	//algos := []string{"Classic", "recur", "fastrec"}
	var results [3][7][3]float64
	var res [3][7][3]float64
	for j := 0; j < 3; j++ {
		i := 0
		for size := 256; size < 1601; size += 224 {
			benchmarks = append(benchmarks, size)
			M1 := grm(size, size)
			M2 := grm(size, size)

			start := time.Now()
			_ = classic(M1, M2)
			restime := float64(time.Since(start))
			results[j][i][0] = restime / 1000000

			startrecur := time.Now()
			_ = recur(M1, M2)
			restimerecur := float64(time.Since(startrecur))
			results[j][i][1] = restimerecur / 1000000

			startfast := time.Now()
			_ = fastrec(M1, M2)
			restimefast := float64(time.Since(startfast))
			results[j][i][2] = restimefast / 1000000
			i++
		}
	}
	for i := 0; i < 7; i++ {
		for j := 0; j < 3; j++ {
			res[0][i][j] = (results[0][i][j] + results[1][i][j] + results[2][i][j]) / 3
			disp1 := (res[0][i][j] - results[0][i][j])
			disp1 *= disp1
			disp2 := (res[0][i][j] - results[1][i][j])
			disp2 *= disp2
			disp3 := (res[0][i][j] - results[2][i][j])
			disp3 *= disp3
			disp := (disp1 + disp2 + disp3) / 3
			res[1][i][j] = math.Pow(disp, 0.5)
			res[2][i][j] = math.Pow(results[0][i][j]*results[1][i][j]*results[2][i][j], 0.333333)
		}
	}
	fmt.Printf("benchmarks / Classic avg / standrt deviation / avg geom / recur avg / standrt deviation / avg geom / fastrec / avg / standrt deviation / avg geom\n")
	for i := 0; i < 7; i++ {
		fmt.Print(256 + 224*i)
		for j := 0; j < 3; j++ {
			fmt.Printf(" | ")
			for k := 0; k < 3; k++ {
				fmt.Printf(" %.2f ", res[k][i][j])
			}
		}
		fmt.Printf("\n")
	}
	/*
		benchmarks / Classic avg / standrt deviation / avg geom    / recur avg / standrt deviation / avg geom    / fastrec avg / standrt deviation / avg geom
		256       |  26.03               2.29           25.93     |  26.31            0.89            26.29     |   23.43                0.48          23.42
		480       |  248.99              9.70           248.80    |  213.14           1.80            213.13    |   170.22               1.76          170.21
		704       |  731.59              19.00          731.34    |  1709.98          5.21            1709.96   |   1206.15              9.28          1206.11
		928       |  2622.75             147.28         2618.52   |  1729.39          21.80           1729.24   |   1208.35              3.72          1208.33
		1152      |  3872.99             138.49         3870.49   |  13651.67         58.65           13651.41  |   8518.64              118.52        8517.74
		1376      |  13479.97            837.39         13453.18  |  13718.76         60.09           13718.49  |   8579.10              150.48        8577.71
		1600      |  18015.59            673.02         18003.04  |  13894.90         182.72          13893.58  |   8524.74              67.84         8524.40

	*/
}
