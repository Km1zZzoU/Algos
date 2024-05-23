package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	p := head

	cash := make([]*int, 100)
	for {
		if p.Next != nil {
			p = p.Next
		} else {
			return nil
		}

	}
}

func main() {
	var l2 ListNode
	var l1 ListNode
	l1.Next = &l2
	l1.Val = 1
	l2.Next = &l1
	l2.Val = 2
	l3 := detectCycle(&l1)
	println(l3.Val)
}
