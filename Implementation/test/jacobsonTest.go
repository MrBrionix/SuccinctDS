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
      ind := rand.Intn(n + 1)
      var (
        ans, ansJacobson interface{}
        str string
      )
      
      str = "Rank"

      ans = ds.Rank(ind)

      ansJacobson = dsJacobson.Rank(ind)

      if showDetails || ans != ansJacobson {
        fmt.Println("Query:")
        fmt.Println(str,"(type)",ind,"(index)")
        fmt.Println("Answer:")
        fmt.Println(ans)
      }
      if showDetails {
        fmt.Println("Jacobson Answer:")
        fmt.Println(ansJacobson)
      }
      if ansJacobson != ans {
        fmt.Println("Sequence:")
        fmt.Println(v)
        fmt.Print("Wrong answer\n","Expected: ",ans,"\nFound: ",ansJacobson,"\n")
        return
      }
    }
  }

  fmt.Println(t*q,"queries have been executed succesfully")
}
