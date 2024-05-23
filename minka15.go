package main

import (
	"fmt"
	"unsafe" // костыль только для адресной арифметики
)

type dynamicarray struct {
	array *int
	len   int
	cap   int
}

func addpointeri(p *int, i int) *int {
	return (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(i)))
}

func malloc(N int) dynamicarray {
	if N < 1 {
		panic("len array can not be < 1")
	}
	array := make([]int, N)
	return dynamicarray{array: &array[0], len: 0, cap: N}
}

func rellocate(dynarr *dynamicarray, double bool) dynamicarray {
	var dynarr2 dynamicarray
	if double {
		array := make([]int, dynarr.cap*2)
		dynarr2 = dynamicarray{&array[0], dynarr.len, dynarr.cap * 2}
	} else {
		array := make([]int, dynarr.cap/2)
		dynarr2 = dynamicarray{&array[0], dynarr.len, dynarr.cap / 2}
	}
	for i := 0; i < dynarr.len; i++ {
		*addpointeri(dynarr2.array, i*8) = *addpointeri(dynarr.array, i*8)
	}
	return dynarr2
}

func push(dynarr *dynamicarray, element int) {
	if dynarr.len == dynarr.cap {
		*dynarr = rellocate(dynarr, true)
	}
	*addpointeri(dynarr.array, dynarr.len*8) = element
	dynarr.len++
}

func del(dynarr *dynamicarray) {
	if dynarr.len == 0 {
		panic("can not be delete element in empty array")
	}
	dynarr.len--
	if dynarr.len*2 < dynarr.cap {
		*dynarr = rellocate(dynarr, false)
	}
}

func take(dynarr dynamicarray, i int) int {
	if i >= dynarr.len {
		panic("index out of range")
	}
	return *addpointeri(dynarr.array, i*8)
}

func printdynarr(dynarr dynamicarray) {
	for i := 0; i < dynarr.len; i++ {
		fmt.Printf("arr[%d] = %d, ", i, take(dynarr, i))
	}
	fmt.Printf("\n")
}

func main() {
	a := malloc(3)
	push(&a, 3)
	printdynarr(a)
	push(&a, 8)
	printdynarr(a)
	push(&a, 9)
	printdynarr(a)
	push(&a, 42)
	printdynarr(a)
	del(&a)
	printdynarr(a)
	del(&a)
	printdynarr(a)
	push(&a, 42)
	printdynarr(a)
	fmt.Println(take(a, 2))
}
