package jacobsonDS

import "math"

type SlowJacobsonDS struct {
  n, dimBlock, dimSuperBlock, numBlock, numSuperBlock int
  blockRanks, superBlockRanks, v []int
  partialRanks [][]int
}

func (ds *SlowJacobsonDS) Build(sequence []bool) {
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
  for _, x := range sequence { //compute the prefix sums
    t = append(t,currT)
    if x {
      currT++
      ds.v = append(ds.v, 1)
    } else {
      ds.v = append(ds.v, 0)
    }
  }
  t = append(t,currT)
  ds.v = append(ds.v, 0)

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

func (ds *SlowJacobsonDS) Rank(i int) int {
  return ds.superBlockRanks[i / ds.dimSuperBlock] + ds.blockRanks[i / ds.dimBlock] + ds.partialRanks[list2index(ds.v[i - (i % ds.dimBlock): i - (i % ds.dimBlock) + ds.dimBlock])][i % ds.dimBlock]
}
