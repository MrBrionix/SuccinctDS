package succinctDS

type SuccinctDS struct {
  succinctRankSelectDS SuccinctRankSelectDS
}

func (ds *SuccinctDS) Build(sequence string) {
  n := len(sequence)
  v := make([]int, n)
  currT := 0

  for i := 0; i < n; i++ {
    if sequence[i] == '(' {
      v[i] = 1
      currT++
    } else {
      v[i] = -1
      currT--
    }
    if currT < 0 {
      panic("invalid balanced parentheses sequence to succinctDS.Build")
    }
  }
  if currT != 0 {
    panic("invalid balanced parentheses sequence to succinctDS.Build")
  }
  
  ds.succinctRankSelectDS.Build(v)
}

func (ds *SuccinctDS) FindExcess(i int) int {
  x := ds.succinctRankSelectDS.Rank(i)
  if ds.succinctRankSelectDS.v[i] == -1 {
    x--
  }
  return x
}

func (ds *SuccinctDS) FindClose(i int) int {
  if ds.succinctRankSelectDS.v[i] == -1 {
    return -1
  }
  
  x := ds.FindExcess(i)
  return ds.succinctRankSelectDS.SelectFirst(i, x) - 1
}

func (ds *SuccinctDS) FindOpen(i int) int {
  if ds.succinctRankSelectDS.v[i] == 1 {
    return -1
  }
  
  x := ds.FindExcess(i)
  return ds.succinctRankSelectDS.SelectLast(i, x)
}

func (ds *SuccinctDS) LeftEnclose(i int) int {
  x := ds.FindExcess(i)
  if x == 0 {
    return -1
  }
  return ds.succinctRankSelectDS.SelectLast(i, x - 1)
}

func (ds *SuccinctDS) RightEnclose(i int) int {
  x := ds.FindExcess(i)
  if x == 0 {
    return -1
  }
  return ds.succinctRankSelectDS.SelectFirst(i, x - 1) - 1
}

func (ds *SuccinctDS) FindEnclose(i int) [2]int {
  return [2]int{ds.LeftEnclose(i), ds.RightEnclose(i)}
}