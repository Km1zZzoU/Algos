package main

import (
	"fmt"
	"unicode"
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

func rellocate(dynarr *dynamicarray) dynamicarray {
	array := make([]int, dynarr.cap*2)
	dynarr2 := dynamicarray{&array[0], dynarr.len, dynarr.cap * 2}
	for i := 0; i < dynarr.len; i++ {
		*addpointeri(dynarr2.array, i*8) = *addpointeri(dynarr.array, i*8)
	}
	return dynarr2
}

func push(dynarr *dynamicarray, element int) {
	if dynarr.len == dynarr.cap {
		*dynarr = rellocate(dynarr)
	}
	*addpointeri(dynarr.array, dynarr.len*8) = element
	dynarr.len++
}

func del(dynarr *dynamicarray) {
	if dynarr.len == 0 {
		panic("can not be delete element in empty array")
	}
	dynarr.len--
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

func pop(dynarr *dynamicarray) int {
	if dynarr.len == 0 {
		panic("can not be pop element in empty array")
	}
	outed := take(*dynarr, dynarr.len-1)
	del(dynarr)
	return outed
}

func peek(dynarr dynamicarray) int {
	if dynarr.len == 0 {
		panic("can not be pop element in empty array")
	}
	outed := take(dynarr, dynarr.len-1)
	return outed
}

//func priority(oper uint8) (int, bool) {
	if oper == '.' {
		return 2, true
	}
	if oper == '^' {
		return 2, true
	}
	if oper == '!' {
		return 3, false
	}
	if oper == '~' {
		return 3, false
	}
	if oper == '*' {
		return 4, false
	}
	if oper == '+' {
		return 5, false
	}
	if oper == '-' {
		return 5, false
	}
	if oper == '<' {
		return 7, false
	}
	if oper == '>' {
		return 7, false
	}
	if oper == '&' { //and
		return 5, false
	}
	if oper == '|' {
		return 6, false
	}
	return 0, false
}

func badhotdog(string string) {
	stack := malloc(1)
	printdynarr(stack)
	for i := range string {
		if unicode.IsDigit((rune)(string[i])) {
			fmt.Printf("%c ", string[i])
		} else if string[i] == '(' {
			push(&stack, (int)(string[i]))
		} else if string[i] == ')' {
			for peek(stack) != '(' {
				tmp := pop(&stack)
				fmt.Printf("%c ", tmp)
			}
			del(&stack)
			tmp := pop(&stack)
			fmt.Printf("%c ", tmp)
		} else {
			if stack.len > 0 {
				for {
					if stack.len > 0 && peek(stack) != '(' {
						p1, lev1 := priority(string[i])
						p2, _ := priority((uint8)(peek(stack)))
						if p1 > p2 || (p1 == p2 && lev1) {
							tmp := pop(&stack)
							fmt.Printf("%c ", tmp)
						} else {
							push(&stack, (int)(string[i]))
							break
						}
					} else {
						push(&stack, (int)(string[i]))
						break
					}
				}
			} else {
				push(&stack, (int)(string[i]))
			}
		}
	}
	for stack.len > 0 {
		tmp := pop(&stack)
		fmt.Printf("%c ", tmp)
	}
}

func main() {
	string2 := "1-5^3*2-3+4"
	badhotdog(string2)
	string3 := "2*(3+4)"
	badhotdog(string3)
	string4 := "2*3+4"
	badhotdog(string4)
	string5 := "!(5-7*4)^6.2+3"
	badhotdog(string5)
}
