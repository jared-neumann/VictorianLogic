package util

import "github.com/skelterjohn/go.matrix"

func Increment(m *matrix.SparseMatrix, i, j int, incr float64) {
  m.Set(i,j,m.Get(i,j)+incr)
}
