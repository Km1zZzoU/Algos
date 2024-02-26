package main

import (
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
	var results [9][3]float64
	i := 0
	for size := 256; size < 2049; size += 224 {
		benchmarks = append(benchmarks, size)
		M1 := grm(size, size)
		M2 := grm(size, size)

		start := time.Now()
		_ = classic(M1, M2)
		restime := float64(time.Since(start))
		results[i][0] = restime / 1000000000

		startrecur := time.Now()
		_ = recur(M1, M2)
		restimerecur := float64(time.Since(startrecur))
		results[i][1] = restimerecur / 1000000000

		startfast := time.Now()
		_ = fastrec(M1, M2)
		restimefast := float64(time.Since(startfast))
		results[i][2] = restimefast / 1000000000

		i++
	}
	/*
	 | Benchmark |  Classic   |   recur    |  fastrec   |
	 |--------------------------------------------------|
	 | 256       |  0.080212  |  0.071176  |  0.062511  |
	 | 480       |  0.760865  |  0.555569  |  0.429305  |
	 | 704       |  1.889631  |  4.163698  |  2.910908  |
	 | 928       |  6.386918  |  4.07202   |  2.873732  |
	 | 1152      |  9.874305  | 32.634299  | 20.953468  |
	 | 1376      | 30.304447  | 32.825395  | 20.341156  |
	 | 1600      | 38.783783  | 33.189794  | 20.639858  |
	 | 1824      | 87.841704  | 32.599347  | 20.365677  |
	 | 2048      | 146.517756 | 34.175543  | 21.571178  |
	*/
}
