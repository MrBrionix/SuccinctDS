package benchmark

import (
  "time"
  "fmt"
  "os"
  "math/rand"
  "succinctDS"
)

func SuccinctRankSelectBenchmark() {

  rand.Seed(0)//time.Now().UnixNano())
  f1, _ := os.Create("data11.txt")
  n := 1024
  
  for n <= 2*67108864 {
    q := 1000000

    
    m1 := 0.0
    
    var ds succinctDS.SuccinctRankSelectDS
    
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
      
    ds.Build(v);
    
    t := 30
    for i := 0; i < t; i++ {
      t1 := 0
      
      ind := make([]int,q)
      
      for j := 0; j < q; j++ {
        ind[j] = rand.Intn(n + 1)
      }
      
      t := time.Now()
      for j := 0; j < q; j++ {
        _ = ds.Rank(ind[j])      
      }
      t1 = int(time.Since(t).Nanoseconds())  
      
      m1 += float64(t1) / float64(q)
      
      fmt.Print(i," ")
    }
    
    m1 = m1 / float64(t)
    fmt.Print("-------------\nResults for ",n,":\n","SuccinctDS: ",m1,"\n")
    
    fmt.Fprint(f1,"(",n,",",m1,")")
    n *= 2
  }
}
