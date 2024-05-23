package main

type FreqStack struct {
	m     map[int]int
	stack []int
}

func Constructor() FreqStack {
	return FreqStack{m: make(map[int]int), stack: make([]int, 0)}
}

func (this *FreqStack) Push(val int) {
	this.m[val]++
	this.stack = append(this.stack, val)
}

func (this *FreqStack) Pop() int {
	maxx := 0
	num := 0
	for _, value := range this.stack {
		count := this.m[value]
		if count >= maxx {
			maxx = count
			num = value
		}
	}
	for i := len(this.stack) - 1; i > -1; i-- {
		if this.stack[i] == num {
			this.stack = append(this.stack[:i], this.stack[i+1:]...)
			break
		}
	}

	this.m[num]--
	return num
}

func main() {
	s := Constructor()
	s.Push(5)
	s.Push(7)
	s.Push(5)
	s.Push(7)
	s.Push(4)
	s.Push(5)

	println(s.Pop())
	println(s.Pop())
	println(s.Pop())
	println(s.Pop())
}
