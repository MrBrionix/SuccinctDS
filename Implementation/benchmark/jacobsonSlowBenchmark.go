package benchmark

import (
  "time"
  "fmt"
  "os"
  "math/rand"
  "jacobsonDS"
)

func JacobsonSlowBenchmark() {

  rand.Seed(0)//time.Now().UnixNano())
  f3, _ := os.Create("data33.txt")
  n := 1024
  
  for n <= 2*67108864 {
    q := 1000000
    
    m3 := 0.0
    
    var dsJacobsonSlow jacobsonDS.SlowJacobsonDS
    
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
      
    dsJacobsonSlow.Build(vb);
    
    t := 30
    for i := 0; i < t; i++ {
      t3 := 0    

      for j := 0; j < q; j++ {
        ind := rand.Intn(n + 1)
        
        t := time.Now()
        _ = dsJacobsonSlow.Rank(ind)
        t3 += int(time.Since(t).Nanoseconds())
      }
      
      m3 += float64(t3) / float64(q)
      
      fmt.Print(i," ")
    }
    
    m3 = m3 / float64(t)
    fmt.Print("-------------\nResults for ",n,":\n","SlowJacobsonDS: ",m3,"\n")
    
    fmt.Fprint(f3,"(",n,",",m3,")")
    n *= 2
  }
}
