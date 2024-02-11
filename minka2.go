package main

import (
	"fmt"
	"strconv"
)

func mulstr(s string, x int) string {
  output := ""
  for i := 0; i < x; i++{
    output += s
  }
  return output
}

func withoutzero(s string) string {
  if s[0] == '0' && len(s) != 1 {
    return withoutzero(s[1:])
  }
  return s
}

func sum(x, y string) string {
  var a, b string
  if len(x) < len(y) {
    z := x
    x = y
    y = z
  }
  xlen := len(x)
  ylen := len(y)
  a = x
  b = mulstr("0", xlen - ylen) + y

  bonuscur := 0
  bonusnext := 0
  output := ""
  for i := xlen - 1; i > -1; i-- {
    bonuscur = bonusnext
    summ := s2i(string(a[i])) + s2i(string(b[i]))
    if summ + bonuscur > 9 {
      bonusnext = 1
      summ -= 10
    } else {
      bonusnext = 0
    }
    output = i2s(summ + bonuscur) + output
  }
  if bonusnext == 1 {
    return "1" + output
  }
  return output
}

func dif(x, y string) string {
  var a, b string
  
  xlen := len(x)
  ylen := len(y)
  
  a = x
  b = ""
  for i := 0; i < (xlen - ylen); i++ {
    b = "0" + b
  }
  b = b + y

  output := ""
  for i := xlen - 1; i > -1; i-- {
    diff := s2i(string(a[i])) - s2i(string(b[i]))
    if diff < 0 {
      minus1 := "1" + mulstr("0", xlen - i)
      a = dif(a, minus1)
      diff += 10
    } 
    output = i2s(diff) + output
  }
  return output
}

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

func mulkaz(x, y string) string {
  xlen := len(x)
  ylen := len(y)
  
  if xlen == 1 && ylen == 1 {
    return i2s(s2i(x) * s2i(y))
  }

  maxlen := max(xlen, ylen)
  x, y = mulstr("0", maxlen - xlen) + x, mulstr("0", maxlen - ylen) + y
  middle := 1
  
  for middle + middle < maxlen {
    middle += middle
  }

  a := x[:maxlen - middle]
  if a == "" {
    a = "0"
  }
  b := x[maxlen - middle:]
  c := y[:maxlen - middle]
  if c == "" {
    c = "0"
  }
  d := y[maxlen - middle:]
  
  step1 := mulkaz(a, c)
  step2 := mulkaz(b, d)
  step3 := mulkaz(sum(a, b), sum(c, d))
  step4 := dif(step3, sum(step2, step1))
  res := sum((step1 + mulstr("0", 2 * middle)), sum(step2, (step4 + mulstr("0", middle))))
  return withoutzero(res)
}

func main() {
  a := "1234567898765432112345678987654321123456789876543212345678987654321123456789876543211234567898765432123456789876543211234567898765432112345678987654321234567898765432112345678987654321123456789876543212345678987654321123456789876543211234567898765432123456789876543211234567898765432112345678987654321234567898765432112345678987654321123456789876543212345678987654321123456789876543211234567898765432123456789876543211234567898765432112345678987654321234567898765432112345678987654321123456789876543212345678987654321123456789876543211234567898765432123456789876543211234567898765432112345678987654321234567898765432112345678987654321123456789876543212345678987654321123456789876543211234567898765432123456789876543211234567898765432112345678987654321234567898765432112345678987654321123456789876543212345678987654321123456789876543211234567898765432123456789876543211234567898765432112345678987654321234567898765432112345678987654321123456789876543212345678987654321123456789876543211234567898765432123456789876543211234567898765432112345678987654321234567898765432112345678987654321123456789876543212345678987654321123456789876543211234567898765432123456789876543211234567898765432112345678987654321234567898765432112345678987654321123456789876543212345678987654321123456789876543211234567898765432123456789876543211234567898765432112345678987654321234567898765432112345678987654321123456789876543212345678987654321123456789876543211234567898765432123456789876543211234567898765432112345678987654321234567898765432112345678987654321123456789876543212345678987654321123456789876543211234567898765432123456789876543211234567898765432112345678987654321234567898765432112345678987654321123456789876543212345678987654321123456789876543211234567898765432123456789876543211234567898765432112345678987654321"
  fmt.Println(mulkaz(a, a))
  fmt.Println(mulkaz("3", "123"))
  fmt.Println(mulkaz("3456", "123"))
  fmt.Println(mulkaz("12345", "567"))
  fmt.Println(mulkaz("1234567", "123456789"))

  //fmt.Println(s2i(a) * s2i(b)) // ПРОВЕРКА
}
