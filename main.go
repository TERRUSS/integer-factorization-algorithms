package main

import "fmt"

func main(){
  var n = 5959
  fmt.Println("Trial Division")
  fmt.Println("v1 ) n=", n, " res=", trial_division(n))
  fmt.Println("v2 ) n=", n, " res=", trial_division2(n))

  fmt.Println("\nFermat's Method")
  fmt.Println("v1 ) n=", n, " res=", fermat_factorization(n))
  //fmt.Println("v2 ) n=", n, " res=", fermat_factorization2(n))
}
