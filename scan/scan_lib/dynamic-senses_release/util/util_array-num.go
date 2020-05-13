package util

import "math"

func Average(array []float64) (avg float64) {
  for _, val := range(array) {
    avg += val
  }
  avg = avg / float64(len(array))
  return
}

/* averages a number of arrays
 * of the same dimensionality (of course)
 */
func AverageArrays(dists [][]float64) (avg []float64) {
  dimensionality := len(dists[0])
  avg = make([]float64, len(dists[0]))
  for _,dist := range(dists) {
    if len(dist) != dimensionality {panic("arrays are not of the same dimensionality!")}
    for idx,val := range(dist) {
      avg[idx] += val
    }
  }
  for idx,_ := range(avg) {
    avg[idx] = avg[idx] / float64(len(dists))
  }
  return avg
}


func Sum(array []int) (sum int) {
  for _,val := range(array) {
    sum += val
  }
  return
}

func SumFloat(array []float64) (sum float64) {
  for _,val := range(array) {
    sum += val
  }
  return
}

func FindMax(array []float64) (maxIdx int, maxVal float64) {
  maxVal = math.Inf(-1)
  maxIdx = -1
  for cIdx,val := range(array) {
    if val > maxVal {
      maxVal = val
      maxIdx = cIdx
    }
  }
  return
}

func FindMaxInt(array []int) (maxIdx, maxVal int) {
  maxVal = int(math.Inf(-1))
  maxIdx = -1
  for idx,val := range(array) {
    if val > maxVal {
      maxVal = val
      maxIdx = idx
    }
  }
  return
}

func Add(array1, array2 [3]float64) ([3]float64) {
  for i:=0 ; i<len(array1) ; i++ {
    array1[i]+=array2[i]
  }
  return array1
}

func Add_scalar(array []float64, scalar float64) {
  for idx,_ := range(array) {
    array[idx] += scalar
  }
}

func Add_weighted(array1, array2 [3]float64, weight float64) ([3]float64) {
  for i:=0 ; i<len(array1) ; i++ {
    array1[i]+= (weight * array2[i])
  }
  return array1
}

func Divide_by_scalar(array [3]float64, div float64) ([3]float64) {
  for i:=0 ; i<len(array) ; i++ {
    array[i] = array[i] / div
  }
  return array
}

func Substract_mean(slice []float64) {
  mean :=0.0
  for _,v := range(slice) {mean += v}
  mean = mean/float64(len(slice))
  for idx,_ := range(slice) {slice[idx] -= mean}
}


func OverlapInt(v1, v2 []int) (overlap int) {
  for _,val1 := range(v1) {
    for _,val2 := range(v2) {
      if val1 == val2 {
	overlap++
      }
    }
  }
  return
}

func RankPositionsByValues(v []float64) (ranking []int) {
  ranking = make([]int, len(v))
  for top := 0 ; top<len(ranking) ; top++ {
    idx, _ := FindMax(v)
    ranking[top] = idx
    v[idx] = math.Inf(-1)
  }
  return
}

func RankPositionsByValues_greaterthreshold(v []float64, threshold float64) (ranking []int) {
  ranking = make([]int, 0)
  for top := 0 ; top<len(v) ; top++ {
    idx, _ := FindMax(v)
    if v[idx] <= threshold {
      return
    }
    ranking = append(ranking, idx)
    v[idx] = math.Inf(-1)
  }
  return
}
