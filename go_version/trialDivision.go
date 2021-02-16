package main

import (
  "time"
)

func trial_division(n int) ([]int, int) {
  defer TimeTrack(time.Now(), "Trial Division")

  var p_list []int
  var p = 2

  var it_count = 1
  for p*p <= n {
    for n % p == 0 {
      p_list = append(p_list, p)
      n /= p
    }
    p += 1

    it_count += 1
  }

  if ( n != 1 ) {
    p_list = append(p_list, n)
  }

  return p_list, it_count
}

func trial_division2(n int) ([]int, int) {
  defer TimeTrack(time.Now(), "Trial Division 2")
  var it_count = 1

  var p_list []int

  for n % 2 == 0 {
    p_list = append(p_list, 2)
    n /= 2

    it_count += 1
  }
  var p = 3

  for p * p <= n {
    if n % p == 0 {
      p_list = append(p_list, p)
      n /= p
    } else {
      p += 2
    }
    it_count += 1
  }
  if n != 1 {
    p_list = append(p_list, n)
  }

  // Only odd number is possible
  return p_list, it_count
}

