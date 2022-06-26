package benchmark

import (
  "time"
  "fmt"
  "os"
  "math/rand"
  "jacobsonDS"
)

func JacobsonSlowBenchmark() {

  rand.Seed(1)//time.Now().UnixNano())
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
      
      ans := 0  

      ind := make([]int,q)
      
      for j := 0; j < q; j++ {
        ind[j] = rand.Intn(n + 1)
      }

      t := time.Now()
      for j := 0; j < q; j++ {
        ans = dsJacobsonSlow.Rank(ind[j])   
        
        if i * q + j == 3476832 {
          fmt.Println(ind[j], ans)
        }
            
      }
      t3 = int(time.Since(t).Nanoseconds())
      
      m3 += float64(t3) / float64(q)
      
      //fmt.Print(i," ")
    }
    
    m3 = m3 / float64(t)
    //fmt.Print("-------------\nResults for ",n,":\n","SlowJacobsonDS: ",m3,"\n")
    
    fmt.Fprint(f3,"(",n,",",m3,")")
    n *= 2
  }
}
