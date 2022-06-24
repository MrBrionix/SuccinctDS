package jacobsonDS

import "math"

const Word_Size = 63

func int2ranklist(x int, size int) (ranklist []int) {
  tmp := []int{}
  for i := 0; i < size; i++ {
    tmp = append(tmp, x % 2)
    x /= 2
  }

  currT := 0
  for i := 0; i < size; i++ {
    ranklist = append(ranklist, currT)
    if tmp[i] == 1 {
      currT++
    }
  }
  return
}

func list2index(list []int) (index int){
  pow := 1
  for _, x := range list {
    index += pow * x
    pow *= 2
  }
  return
}

type JacobsonDS struct {
  n, dimBlock, dimSuperBlock, numBlock, numSuperBlock int
  blockRanks, superBlockRanks, v []int
  partialRanks [][]int
}

func (ds *JacobsonDS) getIndex(l, r int) int {
  if l / Word_Size == r / Word_Size {
    return (ds.v[l / Word_Size] >> (l % Word_Size)) & ((1 << (r-l)) - 1)
  }
  return (ds.v[l / Word_Size] >> (l % Word_Size)) + ((ds.v[r / Word_Size] & ((1 << (r % Word_Size) - 1))) << (Word_Size - (l % Word_Size)))
}

func (ds *JacobsonDS) Build(sequence []bool) {
  ds.n = len(sequence) + 1
  ds.dimBlock = int(math.Ceil(0.5 * math.Log2(float64(ds.n))))
  ds.dimSuperBlock = int(math.Ceil(math.Log2(float64(ds.n)) * math.Log2(float64(ds.n))))
  
  for ds.dimSuperBlock % ds.dimBlock != 0 {
    ds.dimSuperBlock++
  }
  ds.numBlock = (ds.n + ds.dimBlock - 1) / ds.dimBlock
  ds.numSuperBlock = (ds.n + ds.dimSuperBlock - 1) / ds.dimSuperBlock
  t := []int{}
  currT := 0

  ds.v = []int{}
  tmp := []int{}
  for _, x := range sequence { //compute the prefix sums
    t = append(t,currT)
    if x {
      currT++
      tmp = append(tmp, 1)
    } else {
      tmp = append(tmp, 0)
    }
  }
  t = append(t,currT)
  tmp = append(tmp, 0)

  for len(tmp) % Word_Size != 0 {
    tmp = append(tmp, 0)
  }

  for i := 0; i < len(tmp) / Word_Size; i++ {
    ds.v = append(ds.v, list2index(tmp[i * Word_Size : (i + 1) * Word_Size]))
  }

  for i := 0; i < ds.numSuperBlock; i++ { //build first level
    ds.superBlockRanks = append(ds.superBlockRanks, t[i * ds.dimSuperBlock])
  }

  for i := 0; i < ds.numBlock; i++ { //build second level
    ds.blockRanks = append(ds.blockRanks, t[i * ds.dimBlock] - ds.superBlockRanks[(i * ds.dimBlock) / ds.dimSuperBlock])
  }

  for i := 0; i < (1 << ds.dimBlock); i++ {
    ds.partialRanks = append(ds.partialRanks, int2ranklist(i, ds.dimBlock))
  }
}

func (ds *JacobsonDS) Rank(i int) int {
  return ds.superBlockRanks[i / ds.dimSuperBlock] + ds.blockRanks[i / ds.dimBlock] + ds.partialRanks[ds.getIndex(i - i % ds.dimBlock, i - i % ds.dimBlock + ds.dimBlock)][i % ds.dimBlock]
}
