package succinctDS

import (
  "math"
  "segmentTree"  
)

type SuccinctRankSelectDS struct {
  k, n, numBlock, numSuperBlock int // k = blocksize = ceil(logn)
  A, B, T, v []int
  S segmentTree.SegmentTree
}

func (ds *SuccinctRankSelectDS) Build(sequence []int) {
  var s, t []int
  ds.n = len(sequence) + 1
  ds.k = int(math.Ceil(math.Log2(float64(ds.n))))
  ds.numBlock = (ds.n + ds.k - 1) / ds.k //ceil(2n/k)
  ds.numSuperBlock = (ds.numBlock + ds.k - 1) / ds.k // ceil(ceil(2n/k)/k)
  currT := 0
  ds.v = []int{}
  for _, x := range sequence { //compute the prefix sums
    t = append(t, currT)
    currT += x
    ds.v = append(ds.v, x)
  }
  t = append(t, currT)

  ds.T = []int{}
  for i := 0; i < ds.numSuperBlock; i++ { //build first level
    mini := math.MaxInt
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
    mini := math.MaxInt
    for j := i * ds.k; j < (i + 1) * ds.k; j++ {
      if j < len(t) {
        mini =  segmentTree.Min(mini, t[j])
      }
    }
    ds.B = append(ds.B, mini - ds.T[i / ds.k])
    ds.A = append(ds.A, t[i * ds.k] - ds.T[i / ds.k])
  }
}

func (ds *SuccinctRankSelectDS) Rank(i int) int {
  u :=  ds.T[i / (ds.k * ds.k)] + ds.A[i / ds.k]
  for j := ds.k * (i / ds.k); j < i; j++ {
    u += ds.v[j]
  }
  return u
}

func (ds *SuccinctRankSelectDS) SelectLast(i int, u int) int {
  tmp := ds.Rank(i)
  for x := i - 1; x >= ds.k * (i / ds.k); x-- {
    tmp -= ds.v[x]
    if tmp <= u {
      return x
    }
  }

  for x := (i / ds.k) - 1; x >= ds.k * (i / (ds.k * ds.k)); x-- {
    if ds.B[x] + ds.T[i / (ds.k * ds.k)] <= u {
      currt := ds.T[i / (ds.k * ds.k)] + ds.A[x+1]
      for y :=  x * ds.k + ds.k - 1; y >= x * ds.k; y-- {
        currt -= ds.v[y]
        if currt <= u {
          return y
        }
      }
    }
  }

  x := ds.S.FindPrev(i / (ds.k * ds.k), u)

  if x != -1 {
    for y := x * ds.k + ds.k - 1; y >= x * ds.k; y-- {
      if ds.B[y] + ds.T[x] <= u {
        var currt int
        if y < x * ds.k + ds.k - 1 {
          currt = ds.T[x] + ds.A[y+1]
        } else {
          currt = ds.T[x+1]
        }
        for z :=  y * ds.k + ds.k - 1; z >= y * ds.k; z-- {
          currt -= ds.v[z]
          if currt <= u {
            return z
          }
        }
      }
    }
  }
  return -1
}

func (ds *SuccinctRankSelectDS) SelectFirst(i int, u int) int {
  tmp := ds.Rank(i)
  for x := i + 1; x < segmentTree.Min(ds.k * (i / ds.k + 1), ds.n); x++ {
    tmp += ds.v[x - 1]
    if tmp <= u {
      return x
    }
  }

  for x := (i / ds.k) + 1; x < segmentTree.Min(ds.k * (i / (ds.k * ds.k) + 1), ds.numBlock); x++ {
    if ds.B[x] + ds.T[i / (ds.k * ds.k)] <= u {
      currt := ds.T[i / (ds.k * ds.k)] + ds.A[x]
      for y :=  x * ds.k; y < x * ds.k + ds.k; y++ {
        if currt <= u {
          return y
        }
        currt += ds.v[y]
      }
    }
  }

  x := ds.S.FindSucc(i / (ds.k * ds.k), u)
  if x != -1 {
    for y := x * ds.k; y < x * ds.k + ds.k; y++ {
      if ds.B[y] + ds.T[x] <= u {
        currt := ds.T[x] + ds.A[y]
        for z :=  y * ds.k; z < y * ds.k + ds.k; z++ {
          if currt <= u {
            return z
          }
          currt += ds.v[z]
        }
      }
    }
  }
  return -1
}

func (ds *SuccinctRankSelectDS) Select(u int) int {  //classic select for bit sequences
  return ds.SelectLast(ds.n - 1, u)
}
