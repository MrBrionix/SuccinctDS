package slowDS

type SlowDS struct {
  n int
  t []int
  v []bool
}

func (ds *SlowDS) Build(s string) {
  currT := 0
  ds.n = len(s)
  ds.v, ds.t = []bool{}, []int{}
  for _, c := range s { //compute the excess values
    if c == '(' {
      ds.v = append(ds.v,true)
      ds.t = append(ds.t, currT)
      currT++
    } else {
      ds.v = append(ds.v,false)
      currT--
      ds.t = append(ds.t, currT)
    }
    if(currT < 0) {
      panic("invalid balanced parentheses sequence to slowDS.Build")
    }
  }
  if(currT != 0) {
    panic("invalid balanced parentheses sequence to slowDS.Build")
  }
}

func (ds *SlowDS) FindExcess(i int) int {
  return ds.t[i]
}

func (ds *SlowDS) FindFirstExcess(i, val int) int {
  ans := -1
  for j := i + 1; j < ds.n; j++ {
    if ds.t[j] <= val {
      ans = j
      break
    }
  }
  return ans
}

func (ds *SlowDS) FindLastExcess(i, val int) int {
  ans := -1
  for j := i - 1; j >= 0; j-- {
    if ds.t[j] <= val {
      ans = j
      break
    }
  }
  return ans
}

func (ds *SlowDS) FindClose(i int) int {
  if !ds.v[i] {
    return -1
  }
  
  u := ds.FindExcess(i)
  return ds.FindFirstExcess(i, u);
}

func (ds *SlowDS) FindOpen(i int) int {
  if ds.v[i] {
    return -1
  }
  
  u := ds.FindExcess(i)
  return ds.FindLastExcess(i, u);
}

func (ds *SlowDS) LeftEnclose(i int) int {
  u := ds.FindExcess(i)
  return ds.FindLastExcess(i,u-1)
}

func (ds *SlowDS) RightEnclose(i int) int {
  u := ds.FindExcess(i)
  return ds.FindFirstExcess(i,u-1)
}

func (ds *SlowDS) FindEnclose(i int) [2]int {
  return [2]int{ds.LeftEnclose(i), ds.RightEnclose(i)}
}