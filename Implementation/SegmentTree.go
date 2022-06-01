package main

type SegmentTree struct {
  root *SegmentTreeNode
  size int
}

func (st *SegmentTree) Build(v []int) {
  st.size = len(v);
  if(st.size > 0) {
    st.root = &SegmentTreeNode{}
    st.root.Build(v, 0, st.size-1)
  } else {
    st.root = nil
  }
}

func (st *SegmentTree) MinQuery(l, r int) int { // Extra: find the minumum value between l and r (inclusive)
  return st.root.MinQuery(l, r, 0, st.size - 1)
}

func (st *SegmentTree) FindSucc(ind int, val int) int {
  return st.root.FindSucc(ind, val, 0, st.size - 1)
}

func (st *SegmentTree) FindPrev(ind int, val int) int {
  return st.root.FindPrev(ind, val, 0, st.size - 1)
}