package main

import (
  "fmt"
  "strconv"
  "math"
)

func main() {
  a := 12345678987654321
  n := len(strconv.Itoa(a))
  b := 23126
  m := len(strconv.Itoa(b))  
  var c int
  ost := a
  for ost > b {
    for i := 0; i < n - m + 1; i++ {
      number := 0
      ten := int(math.Pow10(n - m - i))
      for (1 + number) * b < ost / ten {
        number++
      }
      ost -= number * b * ten
      c += number * ten 
    }
  }

  fmt.Println(c, ost)  
}
