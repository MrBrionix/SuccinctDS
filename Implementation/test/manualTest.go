package test

import (
  "fmt"
  "succinctDS"
  "slowDS"
)

func ManualTest() {
  var (
    ds succinctDS.SuccinctDS
    dsSlow slowDS.SlowDS
    s string
  )

  fmt.Println("Enter a balanced parentheses sequence:")
  fmt.Scan(&s)
  ds.Build(s)
  dsSlow.Build(s)
  
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
        fmt.Println(ds.FindClose(ind))
      case 1:
        fmt.Println(ds.FindOpen(ind))
      case 2:
        fmt.Println(ds.LeftEnclose(ind))
      case 3:
        fmt.Println(ds.FindEnclose(ind))
    }
    fmt.Print("Slow Answer: ")
    switch tipo {
      case 0:
        fmt.Println(dsSlow.FindClose(ind))
      case 1:
        fmt.Println(dsSlow.FindOpen(ind))
      case 2:
        fmt.Println(dsSlow.LeftEnclose(ind))
      case 3:
        fmt.Println(dsSlow.FindEnclose(ind))
    }
  }
}