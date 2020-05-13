package util

import "math"

func SumExp(s []float64) float64 {
  var sum float64
  for _, val := range s {sum += math.Exp(val)}
  return sum
}

func SoftMax(s []float64) (t []float64) {
  t = make([]float64, len(s))
  denominator := SumExp(s)
  for k,v := range(s) {
    t[k] = v/denominator
  }
  return
}

func SoftMax_single(s []float64, idx int) (float64) {
  return s[idx] / SumExp(s)
}