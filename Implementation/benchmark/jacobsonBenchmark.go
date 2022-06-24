package benchmark

import (
  "time"
  "fmt"
  "os"
  "math/rand"
  "jacobsonDS"
)

func JacobsonBenchmark() {

  rand.Seed(0)//time.Now().UnixNano())  
  f2, _ := os.Create("data22.txt")

  n := 1024
  
  for n <= 2*67108864 {
    q := 1000000

    m2 := 0.0
    
    var dsJacobson jacobsonDS.JacobsonDS
    
    v := make([]int, n)
    vb := make([]bool, n)
    for j := 0; j < n; j++ {
      v[j] = rand.Intn(2)
      if v[j] == 0 {
        vb[j] = false
      } else {
        vb[j] = true
      }
    }
         
    dsJacobson.Build(vb);  
    
    t := 30
    for i := 0; i < t; i++ {
      t2 := 0

      for j := 0; j < q; j++ {
        ind := rand.Intn(n + 1)
        
        t := time.Now()
        _ = dsJacobson.Rank(ind)
        t2 += int(time.Since(t).Nanoseconds())
      }
     
      m2 += float64(t2) / float64(q)
  
      fmt.Print(i," ")
    }
    
    m2 = m2 / float64(t)
    
    fmt.Print("-------------\nResults for ",n,":\n","JacobsonDS: ",m2,"\n")
    
    fmt.Fprint(f2,"(",n,",",m2,")")
    n *= 2
  }
}
