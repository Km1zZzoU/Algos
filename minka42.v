module main

import rand

struct Node {
	key  int
pub:
	val int
	prio int
pub mut:
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

fn (t Treap) print_tree(indent int) {
	if isnil(t.l) && isnil(t.r) {
		println(' '.repeat(indent) + 'Node(key: $t.node.key, prio: $t.node.prio val: $t.node.val sum: $t.node.sum)')
		return
	}

	if !isnil(t.r) {
		t.r.print_tree(indent + 12)
	}

	println(' '.repeat(indent) + 'Node(key: $t.node.key, prio: $t.node.prio val: $t.node.val sum: $t.node.sum)')

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
	}
	mut head := treaps[0]
	for ; head.papa != unsafe { nil }; head = head.papa {}

	return head
}

fn (before &Treap) treap_push_node(new &Treap) {
	unsafe {
		if before.node.key < new.node.key && before.node.prio < new.node.prio {
			new.l = before.r
			if new.l != nil {
				before.node.sum += new.l.node.sum
			}
			before.r = new
			before.node.sum += new.node.val
			new.papa = before
		} else if before.papa == nil {
			new.l = before
			new.node.sum += new.l.node.sum
			before.papa = new
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
		if k > root.node.key {
			r_l, r_r := root.r.split(k)
			if r_l != nil {
				root.r.node.sum -= r_l.node.sum
			}
			if root.r != nil {
				root.node.sum -= root.r.node.sum
			}
			root.r = r_l
			return root, r_r
		} else {
			l_l, l_r := root.l.split(k)
			if l_r != nil {
				root.l.node.sum -= l_r.node.sum
			}
			if root.l != nil {
				root.node.sum -= root.l.node.sum
			}
			root.l = l_r
			return l_l, l_r
		}
	}
}

fn merge(t1 &Treap, t2 &Treap) &Treap {
	unsafe {
		if t1 == nil || t2 == nil {
			return nil
		}
		if t1.node.prio < t2.node.prio {
			t1.node.sum += t2.node.sum
			t1.r = merge(t1.r, t2)
			return t1
		} else {
			t2.node.sum += t1.node.sum
			t2.l = merge(t1, t2.l)
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

fn main() {
	mut arr := [1, 4, 5, 7, 9, 13, 16]
	arr.sort()
	treap := treap_construct(mut arr)
	treap.print_tree(0)
	print(treap.sum(1, 4) == 25)
}
