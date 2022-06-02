package test

import (
  "fmt"
  "succinctDS"
)

func ManualTest() {
  var (
    x succinctDS.SuccinctDS
    s string
  )

  fmt.Println("Enter a balanced parentheses sequence:")
  fmt.Scan(&s)
  x.Build(s)
  
  var q int
  fmt.Println("Enter the number of queries:")
  fmt.Scan(&q)

  for i := 0; i < q; i++ {
    var ind, tipo int

    fmt.Println("Enter type of query (0 = FindClose, 1 = FindOpen, 2 = LeftEnclose, 3 = FindEnclose):")
    fmt.Scan(&tipo)

    fmt.Println("Enter index of query:")
    fmt.Scan(&ind)

    fmt.Print("Answer: ")
    switch tipo {
      case 0:
        fmt.Println(x.FindClose(ind))
      case 1:
        fmt.Println(x.FindOpen(ind))
      case 2:
        fmt.Println(x.LeftEnclose(ind))
      case 3:
        fmt.Println(x.FindEnclose(ind))
    }
  }
}