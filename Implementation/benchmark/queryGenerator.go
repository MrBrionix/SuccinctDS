package benchmark

import (
  "time"
  "fmt"
  "os"
  "math/rand"
)

func GenerateQuery() {
  rand.Seed(time.Now().UnixNano())
  f, _ := os.Create("input.txt")
  n := 1024
  
  for n <= 67108864 {
    q := 1000000
    
    for j := 0; j < n; j++ {
      fmt.Fprint(f,rand.Intn(2)," ")
    }
    
    t := 30
    for i := 0; i < t; i++ {
      for j := 0; j < q; j++ {
        fmt.Fprint(f,rand.Intn(n + 1)," ")
      }
      fmt.Print(i," ")
    } 
    fmt.Print("-------------\nResults for ",n,"\n")
    
    n *= 2
  }
}
