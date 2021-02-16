package main

import (
    "strings"
    "fmt"
    "strconv"

    "flag"

    "github.com/olekukonko/tablewriter"
    "math/big"
)

func arrayToString (T []big.Int) string {
  var ret string

  for i := range T {
    number := T[i]
    //text := strconv.Itoa(number)
    text := big.String(number)
    ret += text
    if i<len(T)-1 { ret+=", "}
  }
  return ret
}

func printResults (n int, data [][]string) {
  fmt.Println("\nFactorization of N =", n)
  tableString := &strings.Builder{}
  table := tablewriter.NewWriter(tableString)
  table.SetHeader([]string{"Method", "Version", "[p, q]", "Summary"})
  table.SetAutoMergeCellsByColumnIndex([]int{0})
  table.SetRowLine(true)
  table.AppendBulk(data)
  table.Render()
  fmt.Println(tableString.String())
}

func parseArgs() (int, string) {
  var nStr = flag.String("N", "32951", "Specify N number to factorize")
  //var nPtr = flag.Int("N", 32951, "Specify N number to factorize")

  methodPtr := flag.String("method", "all", "a string")

  flag.Parse()

  bigInt := &big.Int{}
  n, _ := bigInt.SetString(*nStr, 10)

  var method = *methodPtr
  //var n= *nPtr

  return n, method
}


func main(){
  n, method := parseArgs()
  if method != "all" {
    fmt.Println("using", method,"method")
  }

  var results [][]string
  var res []int
  var summary int

  if method == "all" || method == "trial" {
    res, summary = trial_division(n)
    results = append(results, []string{"Trial Division", "basic", arrayToString(res), strconv.Itoa(summary)})
    res, summary = trial_division2(n)
    results = append(results, []string{"Trial Division", "improvement 1", arrayToString(res), strconv.Itoa(summary)})
  }

  if method == "all" || method == "fermat" {
      // Fermats Methode
    res, summary = fermat_factorization_(n)
    results = append(results, []string{"Fermat", "basic", arrayToString(res), strconv.Itoa(summary)})

    res, summary = fermat_factorization_improvement(n)
    results = append(results, []string{"Fermat", "improvement 1", arrayToString(res), strconv.Itoa(summary)})

    //res, summary = fermat_factorization_trial_improvement(n)
    //results = append(results, []string{"Fermat", "trial improvement", arrayToString(res), strconv.Itoa(summary)})
  }

  printResults(n, results)
}
