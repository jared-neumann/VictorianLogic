package util

import (
  "math/rand"
  "math"
)

func GetArgMaxLog(mult []float64) (idx int) {
  // exp-normalize
  ExpNormalizeAll(mult)
  idx,_ = FindMax(mult)
  return 
}

func Normalize(mult []float64, norm float64) ([]float64) {
  for idx,v := range(mult) {
    mult[idx] = v / norm
  }
  return mult
}


func Normalize2(mult []float64) ([]float64) {
  norm := 0.0
  for _, v := range(mult) {
      norm += v
  }
  for idx,v := range(mult) {
    mult[idx] = v / norm
  }
  return mult
}



func Normalize2d(mult [][]float64) ([][]float64) {
  var norm float64
  for i,_ := range(mult) {
    for j,_ := range(mult[i]) {
      norm += mult[i][j]
    }
  }
  for i,_ := range(mult) {
    for j,v := range(mult[i]) {
      mult[i][j] = v / norm
    }
  }
  return mult
}


func NormalizeAndCumulate(mult []float64, norm float64) ([]float64) {
  mult[0] = mult[0] / norm
  for idx:=1 ; idx<len(mult) ; idx++ {
    mult[idx] = (mult[idx] / norm) + mult[idx-1]
  }
  return mult
}

func Cumulate(mult []float64) ([]float64) {
  for idx:=1 ; idx<len(mult) ; idx++ {
    mult[idx] = mult[idx] + mult[idx-1]
  }
  return mult
}

func Cumulate_copy(m1 []float64) (m2 []float64) {
  m2 = make([]float64, len(m1))
  m2[0] = m1[0]
  for idx:=1 ; idx<len(m1) ; idx++ {
    m2[idx] = m2[idx-1] + m1[idx]
  }
  return m2
}

func GenerateOrderedRandomNumbers(num int) (sample []float64) {
  sample = make([]float64, num)
  r := rand.Float64()
  sample[num-1] = math.Pow(r, (1.0/float64(num)))
  for i:=len(sample)-2 ; i>=0 ; i-- {
    r = rand.Float64()
    sample[i] = sample[i+1]*math.Pow(r, 1.0/float64(i+1))
  }
  return
}

func ScaleProbability(num float64, newBound [2]float64) (float64) {
  return (num) * (newBound[1]-newBound[0]) + newBound[0]
}

