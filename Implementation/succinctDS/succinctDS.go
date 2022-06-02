package succinctDS

import (
  "math"
  "segmentTree"  
)

type SuccinctDS struct {
  k, n, numBlock, numSuperBlock int // k = blocksize = ceil(log2n)
  A, B, T []int
  v []bool
  S segmentTree.SegmentTree
}

func (ds *SuccinctDS) Build(sequence string) {
  var s, t []int
  ds.n = len(sequence) / 2
  ds.k = int(math.Ceil(math.Log2(float64(2 * ds.n))))
  ds.numBlock = (2 * ds.n + ds.k - 1) / ds.k //ceil(2n/k)
  ds.numSuperBlock = (ds.numBlock + ds.k - 1) / ds.k // ceil(ceil(2n/k)/k)
  currT := 0
  for _, c := range sequence { //compute the excess values
    if c == '(' {
      ds.v = append(ds.v,true)
      t = append(t, currT)
      currT++
    } else {
      ds.v = append(ds.v,false)
      currT--
      t = append(t, currT)
    }
    if(currT < 0) {
      panic("invalid balanced parentheses sequence to succinctDS.Build")
    }
  }
  if(currT != 0) {
    panic("invalid balanced parentheses sequence to succinctDS.Build")
  }

  ds.T = []int{}
  for i := 0; i < ds.numSuperBlock; i++ { //build first level
    mini := 2 * ds.n
    for j := i * ds.k * ds.k; j < (i + 1) * ds.k * ds.k; j++ {
      if j < len(t) {
        mini = segmentTree.Min(mini, t[j])
      }
    }
    s = append(s, mini)
    ds.T = append(ds.T, t[i * ds.k * ds.k])
  }
  ds.S.Build(s)

  ds.B = []int{}
  ds.A = []int{}
  for i := 0; i < ds.numBlock; i++ { //build second level
    mini := 2 * ds.n
    for j := i * ds.k; j < (i + 1) * ds.k; j++ {
      if j < len(t) {
        mini =  segmentTree.Min(mini, t[j])
      }
    }
    ds.B = append(ds.B, mini - ds.T[i / ds.k])
    ds.A = append(ds.A, t[i * ds.k] - ds.T[i / ds.k])
  }
}

func (ds *SuccinctDS) FindClose(i int) int {
  if !ds.v[i] {
    return -1
  }
  
  u :=  ds.T[i / (ds.k * ds.k)] + ds.A[i / ds.k] - (i - ds.k * (i / ds.k)) //step 1
  for j := ds.k * (i / ds.k); j < i; j++ {
    if ds.v[j] {
      u = u + 2
    }
  }
  if !ds.v[ds.k * (i / ds.k)] {
    u = u + 1
  }

  tmp := 0 //step 2
  for x := i; x < ds.k * ((i / ds.k) + 1); x++ {
    if ds.v[x] {
      tmp++
    } else {
      tmp--
    }
    if tmp == 0 {
      return x
    }
  }

  for x := (i / ds.k) + 1; x < ds.k * ((i / (ds.k * ds.k)) + 1); x++ { //step 3
    if ds.B[x] + ds.T[i / (ds.k * ds.k)] <= u {
      currt := ds.T[i / (ds.k * ds.k)] + ds.A[x]
      for y := x * ds.k; y < x * ds.k + ds.k; y++ {
        if currt == u {
          return y
        }
        if ds.v[y] {
          currt++
        }
        if !ds.v[y+1] {
          currt--
        }
      }
    }
  }

  x := ds.S.FindSucc(i / (ds.k * ds.k), u) //step 4
  for y := x * ds.k; y < x * ds.k + ds.k; y++ {
    if ds.B[y] + ds.T[x] <= u {
      currt := ds.T[x] + ds.A[y]
      for z :=  y * ds.k; z < y * ds.k + ds.k; z++ {
        if currt == u {
          return z
        }
        if ds.v[z] {
          currt++
        }
        if !ds.v[z + 1] {
          currt--
        }
      }
    }
  }

  return -1 //not found
}

func (ds *SuccinctDS) FindOpen(i int) int {
  if ds.v[i] {
    return -1
  }
  
  u :=  ds.T[i / (ds.k * ds.k)] + ds.A[i / ds.k] - (i - ds.k * (i / ds.k) + 1) //step 1
  for j := ds.k * (i / ds.k); j < i; j++ {
    if ds.v[j] {
      u = u + 2
    }
  }
  if !ds.v[ds.k * (i / ds.k)] {
    u = u + 1
  }

  tmp := 0 //step 2
  for x := i; x >= ds.k * (i / ds.k); x-- {
    if ds.v[x] {
      tmp++
    } else {
      tmp--
    }
    if tmp == 0 {
      return x
    }
  }

  for x := (i / ds.k) - 1; x >= ds.k * (i / (ds.k * ds.k)); x-- { //step 3
    if ds.B[x] + ds.T[i / (ds.k * ds.k)] <= u {
      currt := ds.T[i / (ds.k * ds.k)] + ds.A[x+1]
      for y :=  x * ds.k + ds.k - 1; y >= x * ds.k; y-- {
        if !ds.v[y+1] {
          currt++
        }
        if ds.v[y] {
          currt--
        }
        if currt == u {
          return y
        }
      }
    }
  }

  x := ds.S.FindPrev(i / (ds.k * ds.k), u) //step 4
  for y := x * ds.k + ds.k - 1; y >= x * ds.k; y-- {
    if ds.B[y] + ds.T[x] <= u {
      var currt int
      if y < x * ds.k + ds.k - 1 {
        currt = ds.T[x] + ds.A[y+1]
      } else {
        currt = ds.T[x+1]
      }
      for z :=  y * ds.k + ds.k - 1; z >= y * ds.k; z-- {
        if !ds.v[z + 1] {
          currt++
        }
        if ds.v[z] {
          currt--
        }
        if currt == u {
          return z
        }
      }
    }
  }

  return -1 //not found
}

func (ds *SuccinctDS) LeftEnclose(i int) int {
  u :=  ds.T[i / (ds.k * ds.k)] + ds.A[i / ds.k] - (i - ds.k * (i / ds.k)) //step 1
  for j := ds.k * (i / ds.k); j < i; j++ {
    if ds.v[j] {
      u = u + 2
    }
  }
  if !ds.v[ds.k * (i / ds.k)] {
    u++
  }
  if !ds.v[i] {
    u--
  }
  u--
  if u == -1 {
    return -1 //no solution
  }

  tmp := u + 1 //step 2
  for x := i; x >= ds.k * (i / ds.k); x-- {
    if tmp == u {
      return x
    }
    if !ds.v[x] {
      tmp++
    }
    if ds.v[x-1]{
      tmp--
    }
  }

  for x := (i / ds.k) - 1; x >= ds.k * (i / (ds.k * ds.k)); x-- { //step 3
    if ds.B[x] + ds.T[i / (ds.k * ds.k)] <= u {
      currt := ds.T[i / (ds.k * ds.k)] + ds.A[x+1]
      for y :=  x * ds.k + ds.k - 1; y >= x * ds.k; y-- {
        if !ds.v[y+1] {
          currt++
        }
        if ds.v[y] {
          currt--
        }
        if currt == u {
          return y
        }
      }
    }
  }

  x := ds.S.FindPrev(i / (ds.k * ds.k), u) //step 4
  for y := x * ds.k + ds.k - 1; y >= x * ds.k; y-- {
    if ds.B[y] + ds.T[x] <= u {
      var currt int
      if y < x * ds.k + ds.k - 1 {
        currt = ds.T[x] + ds.A[y+1]
      } else {
        currt = ds.T[x+1]
      }
      for z :=  y * ds.k + ds.k - 1; z >= y * ds.k; z-- {
        if !ds.v[z + 1] {
          currt++
        }
        if ds.v[z] {
          currt--
        }
        if currt == u {
          return z
        }
      }
    }
  }

  return -1 //not found
}

func (ds *SuccinctDS) RightEnclose(i int) int {
  x := ds.LeftEnclose(i)
  if x == -1 {
    return -1
  }
  return ds.FindClose(x)
}

func (ds *SuccinctDS) FindEnclose(i int) [2]int {
  x := ds.LeftEnclose(i)
  if x == -1 {
    return [2]int{-1,-1}
  }
  return [2]int{x,ds.FindClose(x)}
}