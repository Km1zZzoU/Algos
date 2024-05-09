package main

import (
	"fmt"
	"math"
	"math/rand"
)

type bloom struct {
	size uint
	bits []bool
	k    uint
	seed uint
}

func (b bloom) detectbit(ip uint, seed uint) uint {
	var hash uint
	for i := 0; i < 32; i++ {
		hash += uint(math.Pow(float64(seed*ip), float64(i)))
	}
	return hash % b.size
}

func construct(size, k, seed uint) bloom {
	return bloom{size: size, bits: make([]bool, size), k: k, seed: seed}
}

func create(s, seed uint, e float64) bloom {
	b := -(math.Log2(e)) / (math.Pow(math.Ln2, 2))
	k := uint(b * math.Ln2)
	n := uint(b * float64(s))
	return construct(n, k, seed)
}

func (b bloom) insert(ip uint) {
	for i := uint(0); i < b.k; i++ {
		b.bits[b.detectbit(ip, i+b.seed)] = true
	}
}

func (b bloom) lookup(ip uint) bool {
	for i := uint(0); i < b.k; i++ {
		if !b.bits[b.detectbit(ip, i+b.seed)] {
			return false
		}
	}
	return true
}

func main() {
	startseed := uint(rand.Int31())
	filtr := create(20, startseed, 0.05)
	filtr.insert(7635564001)
	filtr.insert(7635564002)
	filtr.insert(7635564003)
	filtr.insert(7635564004)
	filtr.insert(7635564005)
	filtr.insert(7635564006)
	filtr.insert(7635564007)
	filtr.insert(7635564008)
	filtr.insert(7635564009)
	filtr.insert(7635564000)
	filtr.insert(7635564010)
	filtr.insert(7635564011)
	filtr.insert(7635564012)
	filtr.insert(7635564013)
	filtr.insert(7635564014)
	filtr.insert(7635564015)
	filtr.insert(7635564016)
	filtr.insert(7635564017)
	filtr.insert(7635564018)
	filtr.insert(7635564019)
	count := 0
	for i := uint(0); i < 20; i++ {
		if filtr.lookup(7635564000 + i) {
			count++
		}
	}
	fmt.Println(float64(count) / 20)
	count = 0
	for i := uint(0); i < 2<<20; i++ {
		if filtr.lookup(i<<8 + 1337) {
			count++
		}
	}
	fmt.Println(float64(count) / (2 << 20))
}
