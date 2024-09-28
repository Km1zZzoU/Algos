package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 */
type prioq struct {
	array []*ListNode
	len   int
}

func initar(N int) prioq {
	if N < 0 {
		panic("len array can not be < 0")
	}
	array := make([]*ListNode, N)
	return prioq{array, N}
}

func push(priority *prioq, element *ListNode) {
	priority.array = append(priority.array, element)
	priority.len++
}

func swap(priority *prioq, i, j int) {
	tmp := priority.array[i]
	priority.array[i] = priority.array[j]
	priority.array[j] = tmp
}

func fixheadheap(priority *prioq, i int) {
	for i*2+1 < priority.len {
		if i*2+2 >= priority.len || priority.array[i*2+1].Val < priority.array[i*2+2].Val {
			if priority.array[i].Val > priority.array[i*2+1].Val {
				swap(priority, i, i*2+1)
				fixheadheap(priority, i*2+1)
			}
		} else if priority.array[i].Val > priority.array[i*2+2].Val {
			swap(priority, i, i*2+2)
			fixheadheap(priority, i*2+2)
		}
		i++
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	queue := initar(0)
	for i := 0; i < len(lists); i++ {
		list := lists[i]
		if list == nil {
			continue
		}
		push(&queue, list)
	}

	if queue.len == 0 {
		return nil
	}

	h := ListNode{42, nil}
	head := &h
	otv := head

	for i := queue.len / 2; i >= 0; i-- {
		fixheadheap(&queue, i)
	}

	for {
		head.Next = queue.array[0]
		if head.Next != nil {
			head = head.Next
		} else {
			break
		}

		queue.array[0] = queue.array[0].Next
		if queue.array[0] == nil {
			swap(&queue, 0, queue.len-1)
			queue.len--
		}
		fixheadheap(&queue, 0)
	}

	return otv.Next
}

func newNode(val int) *ListNode {
	return &ListNode{Val: val, Next: nil}
}

func createList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := newNode(arr[0])
	current := head

	for i := 1; i < len(arr); i++ {
		current.Next = newNode(arr[i])
		current = current.Next
	}

	return head
}

func printLinkedList(head *ListNode) {
	var builder strings.Builder
	current := head

	for current != nil {
		builder.WriteString(fmt.Sprintf("%d", current.Val))
		if current.Next != nil {
			builder.WriteString(" -> ")
		}
		current = current.Next
	}

	fmt.Println(builder.String())
}

func main() {
	/*
		lists := []*ListNode{
			createList([]int{1}),
			createList([]int{0}),
		}
	*/
	lists := []*ListNode{
		createList([]int{1, 4, 5}),
		createList([]int{1, 3, 4, 7}),
		createList([]int{2, 6}),
		createList([]int{1, 3, 5, 6, 8}),
	}
	head := mergeKLists(lists)
	printLinkedList(head)
}
