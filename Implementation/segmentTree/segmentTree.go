package segmentTree

type SegmentTree struct {
  root *SegmentTreeNode
  size int
}

func (st *SegmentTree) Build(v []int) {
  st.size = len(v)
  if st.size > 0 {
    st.root = &SegmentTreeNode{}
    st.root.Build(v, 0, st.size-1)
  } else {
    st.root = nil
  }
}

func (st *SegmentTree) MinQuery(l, r int) int { // Extra: find the minimum value between l and r (inclusive)
  return st.root.MinQuery(l, r, 0, st.size-1)
}

func (st *SegmentTree) SelectFirst(ind int, val int) int {
  return st.root.SelectFirst(ind, val, 0, st.size-1)
}

func (st *SegmentTree) SelectLast(ind int, val int) int {
  return st.root.SelectLast(ind, val, 0, st.size-1)
}
