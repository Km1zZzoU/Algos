package main

import (
	"fmt"
	"log"
	"strings"
)

type bintree struct {
	val  uint16
	papa *bintree // дерево отец
	hson *binheap // хип сын
}

type binheap struct {
	trees []*bintree // деревья
	count uint8
}

func (bh *binheap) peek_min() uint16 {
	minvalue := bh.trees[0].val
	for i := uint8(0); i < bh.count; i++ {
		minvalue = min(minvalue, bh.trees[i].val)
	}
	return minvalue
}

func swap_value(t1, t2 *bintree) {
	tmp := t1.val
	t1.val = t2.val
	t2.val = tmp
}

func (element *bintree) decrease_key(value uint16) {
	element.val = value
	for element.papa != nil && element.papa.val > element.val {
		swap_value(element, element.papa)
	}
}

func merge_trees(t1, t2 *bintree) *bintree {
	if t2 == nil {
		return t1
	}
	if t1 == nil {
		return t2
	}

	var newbintree *bintree
	if t1.val > t2.val {
		newbintree = &bintree{
			val:  t2.val,
			papa: t2.papa,
			hson: &binheap{
				trees: append(t2.hson.trees, t1),
				count: t2.hson.count + 1,
			},
		}
	} else {
		newbintree = &bintree{
			val:  t1.val,
			papa: t1.papa,
			hson: &binheap{
				trees: append(t1.hson.trees, t2),
				count: t1.hson.count + 1,
			},
		}
	}
	return newbintree
}

func make_bh(size uint8) *binheap {
	newbinheap := &binheap{
		trees: make([]*bintree, size),
		count: size,
	}
	return newbinheap
}

func (h1 *binheap) merge_heaps(h2 *binheap) {
	if h1.count < h2.count {
		tmp := *h1
		*h1 = *h2
		*h2 = tmp
	}
	var sizenewbh uint8

	if h1.count != h2.count && h1.trees[h1.count-1] == nil {
		sizenewbh = h1.count
	} else {
		sizenewbh = h1.count + 1
	}
	newbinheap := make_bh(sizenewbh)

	var mindtree *bintree
	mindtree = nil

	for i := uint8(0); i < h2.count; i++ {
		if mindtree == nil {
			if h1.trees[i] == nil {
				if h2.trees[i] == nil {
					newbinheap.trees[i] = mindtree
				} else {
					newbinheap.trees[i] = h2.trees[i]
				}
			} else {
				if h2.trees[i] == nil {
					newbinheap.trees[i] = h1.trees[i]
				} else {
					newbinheap.trees[i] = nil
					mindtree = merge_trees(h1.trees[i], h2.trees[i])
				}
			}
		} else {
			if h1.trees[i] == nil {
				if h2.trees[i] == nil {
					newbinheap.trees[i] = mindtree
				} else {
					newbinheap.trees[i] = nil
					mindtree = merge_trees(h2.trees[i], mindtree)
				}
			} else {
				if h2.trees[i] == nil {
					newbinheap.trees[i] = nil
					mindtree = merge_trees(h1.trees[i], mindtree)
				} else {
					newbinheap.trees[i] = mindtree
					mindtree = merge_trees(h1.trees[i], h2.trees[i])
				}
			}
		}
	}

	for i := h2.count; i < h1.count; i++ {
		if mindtree == nil {
			newbinheap.trees[i] = h1.trees[i]
		} else {
			if h1.trees[i] == nil {
				newbinheap.trees[i] = mindtree
				mindtree = nil
			} else {
				newbinheap.trees[i] = nil
				mindtree = merge_trees(h1.trees[i], mindtree)
			}
		}
	}
	if mindtree != nil {
		newbinheap.trees[h1.count] = mindtree
	}

	*h1 = *newbinheap
}

func (bh *binheap) insert(element uint16) {
	bh2 := make_bh(1)
	bh2.trees[0] = &bintree{
		val:  element,
		papa: nil,
		hson: &binheap{
			trees: make([]*bintree, 0),
			count: 0,
		},
	}
	bh.merge_heaps(bh2)
}

func (bh *binheap) extract_min() {
	if bh == nil || bh.count == 0 {
		log.Fatal("extract error")
	}
	var (
		minvalue uint16
		indexmin uint8
	)
	for i := uint8(0); i < bh.count; i++ {
		if bh.trees[i] != nil {
			minvalue = bh.trees[i].val
			indexmin = i
			break
		}
	}
	for i := indexmin + 1; i < bh.count; i++ {
		if bh.trees[i] != nil && bh.trees[i].val < minvalue {
			indexmin = i
			minvalue = bh.trees[i].val
		}
	}
	cutbh := bh.trees[indexmin].hson
	bh.trees[indexmin] = nil
	bh.merge_heaps(cutbh)
}

func (bh *binheap) delete(element *bintree) {
	element.decrease_key(0)
	bh.extract_min()
}

func (heap *binheap) visualPrint() {
	for _, tree := range heap.trees {
		heap.visualPrintTree(tree, 0)
	}
}

func (heap *binheap) visualPrintTree(tree *bintree, depth int) {
	if tree == nil {
		return
	}

	fmt.Printf("%s%d\n", strings.Repeat("  ", depth), tree.val)

	if tree.hson != nil {
		for _, subTree := range tree.hson.trees {
			heap.visualPrintTree(subTree, depth+1)
		}
	}
}

func main() {
	binomheap := make_bh(0)
	for i := 30; i > 0; i-- {
		println("\n-=-=-=-=-\n")
		binomheap.insert(uint16(i ^ i ^ i))
		binomheap.visualPrint()
	}
	fmt.Println("----------------------------------------------")
	for i := 29; i > 0; i-- {
		binomheap.extract_min()
		binomheap.visualPrint()
		println("\n-=-=-=-=-\n")
	}
	/*
	 Остальные функции вызываются
	 в ходе инсерта и экстракта,
	 либо являются их
	 упрощенными версиями
	*/
}
