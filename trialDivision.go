package main

func trial_division(n int) []int {
  var a []int
  var f = 2

  for n > 1 {
    if n % f == 0 {
      a = append(a, f)
      n /= f
    } else {
      f += 1
    }
  }

  return a
}

func trial_division2(n int) []int {
  var a []int
  
  for n % 2 == 0 {
    a = append(a, 2)
    n /= 2
  }
  var f = 3

  for f * f <= n {
    if n % f == 0 {
      a = append(a, f)
      n /= f
    } else {
      f += 2
    }
  }
  if n != 1 {
    a = append(a, n)
  }

  // Only odd number is possible
  return a
}

