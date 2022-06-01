package main

func min(x, y int) int {
  if x < y {
    return x 
  } else {
    return y
  }
}

type SegmentTreeNode struct {
  mini int
  left, right *SegmentTreeNode
}

func (node *SegmentTreeNode) Build(v []int, l, r int) {
  if l == r {
    node.mini = v[l]
    return
  }
  node.left = &SegmentTreeNode{}
  node.left.Build(v, l, (l + r) / 2)
  node.right = &SegmentTreeNode{}
  node.right.Build(v, (l + r) / 2 + 1, r)
  node.mini = min(node.left.mini, node.right.mini)
}

func (node *SegmentTreeNode) MinQuery(l , r, _l, _r int) int {
  if _r < l || _l > r || l > r {
    return 0
  }
  if _l >= l && _r <= r {
    return node.mini
  }
  return min(node.left.MinQuery(l, r, _l, (_l + _r) / 2), node.right.MinQuery(l, r, (_l + _r) / 2 + 1, _r))
}

func (node *SegmentTreeNode) FindSucc(ind, val, l, r int) int {
  if ind >= r || node.mini > val {
    return -1
  }
  if l == r {
    return l
  }
  res := node.left.FindSucc(ind, val, l, (l + r) / 2)
  if res != -1 {
    return res
  }
  return node.right.FindSucc(ind, val, (l + r) / 2 + 1, r)
}

func (node *SegmentTreeNode) FindPrev(ind, val, l, r int) int {
  if ind <= l || node.mini > val {
    return -1
  }
  if l == r {
    return l
  }
  res := node.right.FindPrev(ind, val, (l + r) / 2 + 1, r)
  if res != -1 {
    return res
  }
  return node.left.FindPrev(ind, val, l, (l + r) / 2)
}