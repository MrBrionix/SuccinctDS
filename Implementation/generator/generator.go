package generator

import "math/rand"

func SetSeed(x int64) {
  rand.Seed(x)
}

func RandSeq(length int) (res string) {
  var x, y, n int64
  n, x, y, res = int64(length), 0, 0, ""
  for x != n || y != n {
    var a, b int64
    a = (n - x) * (x - y + 2)
    b = (x - y) * (n - y + 1)
    rn := rand.Int63n(a + b)
    if rn < a {
      res += "("
      x++
    } else {
      res += ")"
      y++
    }
  }
  return
}