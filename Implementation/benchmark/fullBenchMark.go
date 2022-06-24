package benchmark

import (
  "time"
  "fmt"
  "os"
  "math/rand"
  "succinctDS"
  "jacobsonDS"
)

func FullBenchmark() {

  rand.Seed(0)//time.Now().UnixNano())
  f1, _ := os.Create("data11.txt")
  f2, _ := os.Create("data22.txt")
  f3, _ := os.Create("data33.txt")
  n := 1024
  
  for n <= 67108864 {
    q := 1000000

    
    m1 := 0.0
    m2 := 0.0
    m3 := 0.0
    
    var (
      ds succinctDS.SuccinctRankSelectDS
      dsJacobson jacobsonDS.JacobsonDS
      dsJacobsonSlow jacobsonDS.SlowJacobsonDS
    )
    
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
    dsJacobson.Build(vb);
    dsJacobsonSlow.Build(vb);
    
    t := 30
    for i := 0; i < t; i++ {
      t1 := 0
      t2 := 0
      t3 := 0
      
      tmp := make([]int,q)
      for j := 0; j < q; j++ {
        tmp[j] = rand.Intn(n + 1)
      }

      for j := 0; j < q; j++ {
        ind := tmp[j]     
        t := time.Now()
        _ = ds.Rank(ind)
        t1 += int(time.Since(t).Nanoseconds())
      }
      
      for j := 0; j < q; j++ {
        ind := tmp[j]
        t := time.Now()
        _ = dsJacobson.Rank(ind)
        t2 += int(time.Since(t).Nanoseconds())
      }
      
      for j := 0; j < q; j++ {  
        ind := tmp[j]
        t := time.Now()
        _ = dsJacobsonSlow.Rank(ind)
        t3 += int(time.Since(t).Nanoseconds())
      }  
      
      
      m1 += float64(t1) / float64(q)
      m2 += float64(t2) / float64(q)
      m3 += float64(t3) / float64(q)
      
      fmt.Print(i," ")
      //fmt.Print("-------------\nIteration ",i,", time elapsed:\n","SuccinctDS: ",float64(t1) / float64(q),"\nJacobsonDS: ",float64(t2) / float64(q),"\nSlowJacobsonDS: ",float64(t3) / float64(q),"\n")
    }
    
    m1 = m1 / float64(t)
    m2 = m2 / float64(t)
    m3 = m3 / float64(t)
    fmt.Print("-------------\nResults for ",n,":\n","SuccinctDS: ",m1,"\nJacobsonDS: ",m2,"\nSlowJacobsonDS: ",m3,"\n")
    
    fmt.Fprint(f1,"(",n,",",m1,")")
    fmt.Fprint(f2,"(",n,",",m2,")")
    fmt.Fprint(f3,"(",n,",",m3,")")
    n *= 2
  }
}
