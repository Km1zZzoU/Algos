package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
  a := 12345678987654321
  n := len(strconv.Itoa(a))
  b := 23126
  m := len(strconv.Itoa(b))
  var c int
  ost := a
  for i := 0; i < n - m + 1; i++ {
    number := 0
    ten := int(math.Pow10(n - m - i))

    for ost - ten * b > 0 {
      ost -= ten * b
      number++
    }

    c += number * ten
  }

  fmt.Println(c, ost)
}
