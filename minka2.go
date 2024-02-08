package main

import (
  "fmt"
  "strconv"
  "math"
)

func i2s(x int) string {
  return strconv.Itoa(x)
}

func s2i(str string) int {
  x, err := strconv.Atoi(str)
  if err == nil {
    return x
  }
  return 0
}

func mulkaz(x, y int) int {
  xlen := len(i2s(x))
  ylen := len(i2s(y))
  
  if xlen == 1 && ylen == 1 {
    return x * y
  }
  
  minlen := min(xlen, ylen)

  var middle int = minlen / 2
  if xlen + ylen == 3 {
    middle = 1
  }

  a := s2i(i2s(x)[:xlen - middle])
  b := s2i(i2s(x)[xlen - middle:])
  c := s2i(i2s(y)[:ylen - middle])
  d := s2i(i2s(y)[ylen - middle:])
  
  step1 := mulkaz(a, c)
  step2 := mulkaz(b, d)
  step3 := mulkaz(a + b, c + d)
  step4 := step3 - step2 - step1

  res := step1 * int(math.Pow10(xlen)) + step2 + step4 * int(math.Pow10(middle))
  return res
}

func main() {
  fmt.Println(mulkaz(1234, 5678))
}
