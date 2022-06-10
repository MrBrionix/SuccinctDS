package test

import (
  "time"
  "fmt"
  "math/rand"
  "succinctDS"
  "slowDS"
)

func RankSelectDS(showDetails bool, slowCheck bool) {
  var t, n, m, q int
  fmt.Println("Enter t (number of tests), n (sequence size), m (value size) and q (queries size)")
  fmt.Scan(&t, &n, &m, &q)

  rand.Seed(time.Now().UnixNano())
  for i := 0; i < t; i++ {
    var (
      ds succinctDS.SuccinctRankSelectDS
      dsSlow slowDS.SlowRankSelectDS
    )
  
    v := make([]int, n)
    for j := 0; j < n; j++ {
      v[j] = rand.Intn(2 * m - 1) - (m - 1) //random number in [-m+1, m-1]
    }
    if showDetails {
      fmt.Println("Sequence:")
      fmt.Println(v)
    }

    ds.Build(v);
    dsSlow.Build(v);

    for j := 0; j < q; j++ {
      mode := rand.Intn(4)
      ind := rand.Intn(n)
      val := 0
      var (
        ans, ansSlow interface{}
        str string
      )
      switch mode {
        case 0:
          str = "Rank"
          ans = ds.Rank(ind)
        case 1:
          str = "SelectLast"
          val = rand.Intn(n * (2 * m - 1)) - n * (m - 1)
          ans = ds.SelectLast(ind, val)
        case 2:
          str = "SelectFirst"
          val = rand.Intn(n * (2 * m - 1)) - n * (m - 1)
          ans = ds.SelectFirst(ind, val)
        case 3:
          str = "Select"
          val = rand.Intn(n * (2 * m - 1)) - n * (m - 1)
          ans = ds.Select(val)
      }
      if slowCheck {
        switch mode {
          case 0:
            ansSlow = dsSlow.Rank(ind)
          case 1:
            ansSlow = dsSlow.SelectLast(ind, val)
          case 2:
            ansSlow = dsSlow.SelectFirst(ind, val)
          case 3:
            ansSlow = dsSlow.Select(val)
        }
      }
      if showDetails || (slowCheck && ans != ansSlow) {
        fmt.Println("Query:")
        fmt.Println(str,"(type)",ind,"(index)",val,"(val)")
        fmt.Println("Answer:")
        fmt.Println(ans)
      }
      if showDetails && slowCheck {
        fmt.Println("Slow Answer:")
        fmt.Println(ansSlow)
      }
      if slowCheck && ansSlow != ans {
        fmt.Println("Sequence:")
        fmt.Println(v)
        fmt.Print("Wrong answer\n","Expected: ",ansSlow,"\nFound: ",ans,"\n")
        return
      }
    }
  }

  fmt.Println(t*q,"queries have been executed succesfully")
}