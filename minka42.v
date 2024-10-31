module main

import rand

struct Node {
	key  int
pub:
	val int
	prio int
pub mut:
	c   int
	sum int
}

@[heap]
struct Treap {
pub mut:
  node Node
	papa &Treap = unsafe { nil }
	l    &Treap = unsafe { nil }
	r    &Treap = unsafe { nil }
}

fn (t &Treap) print_tree(indent int) {
	if t == unsafe { nil } {
		println("nil")
		return
	}
	if isnil(t.l) && isnil(t.r) {
		println(' '.repeat(indent) + 'Node(c: $t.node.c, prio: $t.node.prio val: $t.node.val sum: $t.node.sum)')
		return
	}

	if !isnil(t.r) {
		t.r.print_tree(indent + 12)
	}

	println(' '.repeat(indent) + 'Node(c: $t.node.c, prio: $t.node.prio val: $t.node.val sum: $t.node.sum)')

	if !isnil(t.l) {
		t.l.print_tree(indent + 12)
	}
}

pub fn treap_construct(mut values []int) &Treap {
	if values.len == 0 {
		panic("values is empty")
	}

	mut nodes := []Node{}
	for i in 0..values.len {
		nodes << Node{
			key:  i
			prio: rand.int()
			sum:  values[i]
			val:  values[i]
			c:    1
		}
	}

	mut treaps := []&Treap{}
	for i in 0..values.len {
		treaps << &Treap{
			node: nodes[i]
			papa: unsafe { nil }
			l:    unsafe { nil }
			r:    unsafe { nil }
		}
	}
	for i in 1..values.len {
		treaps[i-1].treap_push_node(treaps[i])

		mut head := treaps[0]
		for ; head.papa != unsafe { nil }; head = head.papa {}
		head.print_tree(0)
		println("-----------====----------")
	}
	mut head := treaps[0]
	for ; head.papa != unsafe { nil }; head = head.papa {}

	return head
}

fn (treap &Treap) update_c_sum_rec(n int) {
	unsafe {
		for ;treap != nil; treap = treap.papa {
			treap.node.sum += n
			treap.node.c++
		}
	}
}

fn (treap &Treap) update() {
	unsafe {
		treap.node.c = 1
		treap.node.sum = treap.node.val
		if treap.l != nil {
			treap.node.c += treap.l.node.c
			treap.node.sum += treap.l.node.sum
		}
		if treap.r != nil {
			treap.node.c += treap.r.node.c
			treap.node.sum += treap.r.node.sum
		}
	}
}

fn (before &Treap) treap_push_node(new &Treap) {
	unsafe {
		if before.node.key < new.node.key && before.node.prio < new.node.prio {
			new.l = before.r
			if new.l != nil {
				new.node.sum += new.l.node.sum
			}
			before.r = new
			before.update_c_sum_rec(new.node.val)
			new.papa = before
		} else if before.papa == nil {
			new.l = before
			new.node.sum += new.l.node.sum
			before.papa = new
			new.update()
		} else {
			before.papa.treap_push_node(new)
		}
	}
}

fn (root &Treap) split(k int) (&Treap, &Treap) {
	unsafe {
		if root == nil {
			return nil, nil
		}
		if root.l == nil || k > root.l.node.c {
			lc := 0
			if root.l != nil {
				lc = root.l.node.c
			}
			r_r := root.r
			r_l := root.r
			if k - lc - 1 > 0 {
				r_l, r_r = root.r.split(k - lc - 1)
				root.r = r_l
			} else {
				root.r = nil
			}
			if r_r != nil {
				root.update()
			}
			return root, r_r
		} else {
			l_l, l_r := root.l.split(k)
			root.l = l_r
			root.update()
			return l_l, root
		}
	}
}

fn merge(t1 &Treap, t2 &Treap) &Treap {
	unsafe {

		if t1 == nil {
			return t2
		}

		if t2 == nil {
			return t1
		}

		if t1.node.prio < t2.node.prio {
			t1.r = merge(t1.r, t2)
			t1.update()
			return t1
		} else {
			t2.l = merge(t1, t2.l)
			t2.update()
			return t2
		}
	}
}

pub fn (root &Treap) sum(from int, to int) int {
	l, r := root.split(from)
	rl, rr := r.split(to - from + 1)
	res := 0
	unsafe {
		if rl != nil {
			res = rl.node.sum
		}
		root = merge(l, merge(rl, rr))
	}
	return res
}

fn sum(arr []int, from int, to int) int {
	mut sum := 0
	for i in from..(to+1) {
		sum += arr[i]
	}
	return sum
}
fn main() {
	mut arr := [1, 4, 5, 7, 9, 13, 16]
	arr.sort()
	treap := treap_construct(mut arr)
	println("complete construct")
	treap.print_tree(0)
	for i in 1..4 {
		for j in i..5 {
			print("$i, $j [${treap.sum(i, j)}, ${sum(arr, i, j)}] ")
			println(treap.sum(i, j) == sum(arr, i, j))
		}
	}
}
