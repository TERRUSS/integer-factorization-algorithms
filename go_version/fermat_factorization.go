package main

import (
  "math/big"
  //"math"
  "time"
  "fmt"
)


func fermat_factorization_(input big.Int) ([]int, int) {
  defer TimeTrack(time.Now(), "Fermat's method basic")
  var it_count = 1

  //var n = float64(input)
  var n = new(big.Float).SetInt(&input)
  var a = new(big.Float).Sqrt(n)
  bint, _ := a.Int(nil)
  a.SetInt(bint)

  fmt.Println(n, a)

  tmp := new(big.Float)

  if tmp.Mul(a, a) == n {
    return []int{a.Int(nil), a.Int(nil)}, it_count
  }

  b2 := new(big.Float).Mul(a,a)
  b2 = b2.Sub(b2, n)

  var b = new(big.Float).Sqrt(b2)
  bint, _ = b2.Int(nil)
  b.setInt(bint)

  for tmp.Mul(b,b) != b2 {
    a = a + 1
    b2 = a*a - n
    b = big.Floor(big.Sqrt(b2))
    it_count++
  }

  return []int{int(a-b), int(a+b)}, it_count
}


func fermat_factorization_improvement(input int) ([]int, int) {
  defer TimeTrack(time.Now(), "Fermat's method improvement 1")
  var n = float64(input)

  var sqrtN = big.Floor(big.Sqrt(n))
  var sqrtN2 = sqrtN*sqrtN

  var PCS = big.Floor(big.Sqrt(big.Abs(sqrtN2 - n)))

  var it_count = 1
  for int(n) % int(PCS + sqrtN) != 0 {

    sqrtN++
    sqrtN2 = sqrtN*sqrtN

    PCS = big.Floor(big.Sqrt(big.Abs(sqrtN2 - n)))

    it_count +=1
  }
  return []int{int(sqrtN-PCS), int(sqrtN+PCS)}, it_count
}

