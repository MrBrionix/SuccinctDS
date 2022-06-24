package benchmark

import (
  "time"
  "fmt"
  "os"
  "succinctDS"
  "jacobsonDS"
)

func JacobsonBenchmarkFscanf() {

  
  f1, _ := os.Create("data1.txt")
  f2, _ := os.Create("data2.txt")
  f3, _ := os.Create("data3.txt")
  f, _ := os.Open("input.txt")
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
      fmt.Fscanf(f,"%d",&v[j])
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

      for j := 0; j < q; j++ {
        var ind int
        fmt.Fscanf(f,"%d",ind)
        
        t := time.Now()
        ans1 := ds.Rank(ind)
        t1 += int(time.Since(t).Nanoseconds())

        t = time.Now()
        ans2 := dsJacobson.Rank(ind)
        t2 += int(time.Since(t).Nanoseconds())
        
        t = time.Now()
        ans3 := dsJacobsonSlow.Rank(ind)
        t3 += int(time.Since(t).Nanoseconds())
        
        
        if ans1 != ans2 || ans1 != ans3 {
          panic("wrong answer")
        }
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
