package test

import (
  "time"
  "fmt"
  "math/rand"
  "generator"
  "succinctDS"
)

func AutomatedTest(showDetails bool) {
  var t, n, q int
  fmt.Println("Enter t (number of tests), n (sequence size) and q (queries size)")
  fmt.Scan(&t, &n, &q) // t = number of tests, n = sequence size, q = queries size
  generator.SetSeed(time.Now().UnixNano())

  for i := 0; i < t; i++ {
    var ds succinctDS.SuccinctDS
    s := generator.RandSeq(n)
    ds.Build(s)
    if showDetails {
      fmt.Println("Sequence:")
      fmt.Println(s)
    }
    
    for j := 0; j < q; j++ {
      mode := rand.Intn(4)
      ind := rand.Intn(2 * n)
      var (
        ans interface{}
        str string
      )
      
      switch mode {
        case 0:
          str = "FindClose"
          ans = ds.FindClose(ind)
        case 1:
          str = "FindOpen"
          ans = ds.FindOpen(ind)
        case 2:
          str = "LeftEnclose"
          ans = ds.LeftEnclose(ind)
        case 3:
          str = "FindEnclose"
          ans = ds.FindEnclose(ind)
      }
      if showDetails {
        fmt.Println("Query:")
        fmt.Println(str,"(type)",ind,"(index)")
        fmt.Println("Answer:")
        fmt.Println(ans)
      }
    }
  }

  fmt.Println(t*q,"queries have been executed succesfully")
}