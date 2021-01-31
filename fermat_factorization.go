package main

import "math"

func fermat_factorization(n int) []int {
  if n % 2 == 0 {
    return []int{n/2, 2}
  }

  var a = math.Ceil( math.Sqrt( float64(n)  ) )

  if a*a == float64(n) {
    return []int{int(a), int(a)}
  }

  var b = 0
  for true {
    var b2 = a*a - float64(n)
    b = int( math.Sqrt(b2) )

    if float64(b*b) == b2 {
      break
    } else {
      a += 1
    }
  }

  return []int{int(a)-b, int(a)+b}
}
