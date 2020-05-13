package util

import (
       "math"
       "fmt"
       "github.com/gonum/floats"
)

func Exp(numbers []float64) ([]float64) {
  for i,_ := range(numbers) {
    numbers[i] = math.Exp(numbers[i])
  }
  return numbers
}


func Exp_copy(v1 []float64) (v2 []float64) {
    v2 = make([]float64, len(v1))
    for i,_ := range(v1) {
        v2[i] = math.Exp(v1[i])
    }
    return v2
}

func expSum(numbers []float64, b float64) (z float64) {
  for _,num := range(numbers) {
    z += math.Exp(num-b)
  }
  if math.IsInf(z,-1) || math.IsInf(z,1) || z==0.0 {
    fmt.Println(z, " ")
    panic("infinity or zero in expSum!")
  }
  return
}


func ExpSum(numbers []float64) (max, z float64) {
  _,max = FindMax(numbers) 
  for _,num := range(numbers) {
    z += math.Exp(num-max)
  }
  if math.IsInf(z,-1) || math.IsInf(z,1) || z==0.0 {
    fmt.Println(z, " ")
    panic("infinity or zero in expSum!")
  }
  return
}


func ExpNormalize(target, b, z float64) float64 {
  tmp := math.Exp(target-b) / z
  if math.IsInf(tmp,-1) || math.IsInf(tmp,1) {
    fmt.Printf("raw %.3f , exp %f , scaled %f , normalized %f , norm_const %.3f\n", target, math.Exp(target), math.Exp(target-b) , tmp , z)
    panic("infinity or zero in ExpNormalize!")
  }
  return tmp
}

func ExpNormalizeAll(numbers []float64) ([]float64) {
  _,b := FindMax(numbers) 
  z := expSum(numbers,b)
  for idx, val := range(numbers) {
    numbers[idx] = ExpNormalize(val, b, z)
  }
  return numbers
}

func ExpNormalizeLog(numbers []float64) () {
  z := floats.LogSumExp(numbers)
  for idx, _ := range(numbers) {
    numbers[idx] = numbers[idx] - z
  }
  return
}

/* takes log weights and log values and returns the log 
 * of the weighted average 
 * input: values in LOG space
 *        NORMALIZED weights in LOG space
 * output: weighted LOG average */
func LogAverageExp(weights,values []float64) (avg float64) {
  if len(weights) != len(values) {
    panic("number of weights and values differ!")
  }
  copy_weights := make([]float64, len(weights))
  copy(copy_weights, weights) 
  copy_weights = Exp(copy_weights)
  m := floats.Max(values)
  for idx,_ := range(values) {
    avg += copy_weights[idx] * math.Exp(values[idx]-m)
  }
  return m+math.Log(avg)
}
