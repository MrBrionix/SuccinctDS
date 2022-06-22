package test

import (
  "time"
  "fmt"
  "math/rand"
  "succinctDS"
  "jacobsonDS"
)

func JacobsonTest(showDetails bool) {
  var t, n, q int
  fmt.Println("Enter t (number of tests), n (sequence size) and q (queries size)")
  fmt.Scan(&t, &n, &q)

  rand.Seed(time.Now().UnixNano())
  t1 := 0
  t2 := 0
  
  for i := 0; i < t; i++ {
    var (
      ds succinctDS.SuccinctRankSelectDS
      dsJacobson jacobsonDS.JacobsonDS
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
    if showDetails {
      fmt.Println("Sequence:")
      fmt.Println(v)
    }

    ds.Build(v);
    dsJacobson.Build(vb);

    for j := 0; j < q; j++ {
      ind := rand.Intn(n)
      var (
        ans, ansJacobson interface{}
        str string
      )
      
      str = "Rank"

      t := time.Now()
      ans = ds.Rank(ind)

      //fmt.Println(time.Since(t).Nanoseconds())
      t1 += int(time.Since(t).Nanoseconds())

      t = time.Now()
      ansJacobson = dsJacobson.Rank(ind)

      //fmt.Println(time.Since(t).Nanoseconds())
      t2 += int(time.Since(t).Nanoseconds())

      if showDetails || ans != ansJacobson {
        fmt.Println("Query:")
        fmt.Println(str,"(type)",ind,"(index)")
        fmt.Println("Answer:")
        fmt.Println(ans)
      }
      if showDetails {
        fmt.Println("Slow Answer:")
        fmt.Println(ansJacobson)
      }
      if ansJacobson != ans {
        fmt.Println("Sequence:")
        fmt.Println(v)
        fmt.Print("Wrong answer\n","Expected: ",ansJacobson,"\nFound: ",ans,"\n")
        return
      }
    }
  }

  fmt.Println(t*q,"queries have been executed succesfully")
  fmt.Println("Time elapsed:",t1,t2)
}
