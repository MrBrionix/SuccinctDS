package slowDS

type SlowRankSelectDS struct{
  n int
  t, v []int
}

func (ds *SlowRankSelectDS) Build(s []int) {
  ds.n = len(s) + 1
  ds.t = []int{}
  ds.v = []int{}
  currT := 0
  for _, x := range s {
    ds.t = append(ds.t, currT)
    ds.v = append(ds.v, x)
    currT += x
  }
  ds.t = append(ds.t, currT)
}

func (ds *SlowRankSelectDS) Rank(i int) int {
  return ds.t[i]
}

func (ds *SlowRankSelectDS) SelectLast(i int, u int) int {
  for x := i - 1; x >= 0; x-- {
    if ds.t[x] <= u {
      return x
    }
  }
  return -1
}

func (ds *SlowRankSelectDS) SelectFirst(i int, u int) int {
  for x := i + 1; x < ds.n; x++ {
    if ds.t[x] <= u {
      return x
    }
  }
  return -1
}

func (ds *SlowRankSelectDS) Select(u int) int {
  return ds.SelectLast(ds.n - 1, u)
}