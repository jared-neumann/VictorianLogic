package util

import (
        "sort"
	"math/rand"
)

type KeyValuePair struct {
  Key string
  Value int
}

type KeyValuePairs []*KeyValuePair

func (pairs *KeyValuePairs)Swap(i,j int)        {(*pairs)[i], (*pairs)[j] = (*pairs)[j], (*pairs)[i]}
func (pairs *KeyValuePairs)Less(i,j int) (bool) {return (*pairs)[i].Value < (*pairs)[j].Value}
func (pairs *KeyValuePairs)Len() (int)          {return len(*pairs)}

func SortKeysByValues(input map[string]int) (pairs KeyValuePairs) {
  pairs = make(KeyValuePairs, len(input))
  var idx int
  for k,v := range(input) {
    pairs[idx] = &KeyValuePair{k, v}
    idx++
  }
  sort.Sort(&pairs)
  return
}





type KeyStringValueFloatPair struct {
  Key string
  Value float64
}

type KeyStringValueFloatPairs []*KeyStringValueFloatPair

func (pairs *KeyStringValueFloatPairs)Swap(i,j int)        {(*pairs)[i], (*pairs)[j] = (*pairs)[j], (*pairs)[i]}
func (pairs *KeyStringValueFloatPairs)Less(i,j int) (bool) {return (*pairs)[i].Value < (*pairs)[j].Value}
func (pairs *KeyStringValueFloatPairs)Len() (int)          {return len(*pairs)}

func SortStringKeysByFloatValues(input map[string]float64) (pairs KeyStringValueFloatPairs) {
  pairs = make(KeyStringValueFloatPairs, len(input))
  var idx int
  for k,v := range(input) {
    pairs[idx] = &KeyStringValueFloatPair{k, v}
    idx++
  }
  sort.Sort(&pairs)
  return
}










type KeyIntlistValueFloatPair struct {
  Key [2]int
  Value float64
}

type KeyIntlistValueFloatPairs []*KeyIntlistValueFloatPair

func (pairs *KeyIntlistValueFloatPairs)Swap(i,j int)        {(*pairs)[i], (*pairs)[j] = (*pairs)[j], (*pairs)[i]}
func (pairs *KeyIntlistValueFloatPairs)Less(i,j int) (bool) {return (*pairs)[i].Value < (*pairs)[j].Value}
func (pairs *KeyIntlistValueFloatPairs)Len() (int)          {return len(*pairs)}

func SortIntlistKeysByFloatValues(input map[[2]int]float64) (pairs KeyIntlistValueFloatPairs) {
  pairs = make(KeyIntlistValueFloatPairs, len(input))
  var idx int
  for k,v := range(input) {
    pairs[idx] = &KeyIntlistValueFloatPair{k, v}
    idx++
  }
  sort.Sort(&pairs)
  return
}



/* sort by key */
type KeyIntValueFloatPair struct {
  Key int
  Value float64
}

/* sort by value */
type KeyIntValueFloatPair2 struct {
  Key int
  Value float64
}

type KeyIntValueFloatPairs []*KeyIntValueFloatPair

func (pairs *KeyIntValueFloatPairs)Swap(i,j int)        {(*pairs)[i], (*pairs)[j] = (*pairs)[j], (*pairs)[i]}
func (pairs *KeyIntValueFloatPairs)Less(i,j int) (bool) {return (*pairs)[i].Value > (*pairs)[j].Value}
func (pairs *KeyIntValueFloatPairs)Len() (int)          {return len(*pairs)}

type KeyIntValueFloatPairs2 []*KeyIntValueFloatPair2

func (pairs *KeyIntValueFloatPairs2)Swap(i,j int)        {(*pairs)[i], (*pairs)[j] = (*pairs)[j], (*pairs)[i]}
func (pairs *KeyIntValueFloatPairs2)Less(i,j int) (bool) {return (*pairs)[i].Key < (*pairs)[j].Key}
func (pairs *KeyIntValueFloatPairs2)Len() (int)          {return len(*pairs)}


func SortIntKeysByFloatValues(input map[int]float64) (pairs KeyIntValueFloatPairs) {
  pairs = make(KeyIntValueFloatPairs, len(input))
  var idx int
  for k,v := range(input) {
    pairs[idx] = &KeyIntValueFloatPair{k, v}
    idx++
  }
  sort.Sort(&pairs)
  return
}


func SortIntKeys_float(input map[int]float64) (pairs KeyIntValueFloatPairs2) {
  pairs = make(KeyIntValueFloatPairs2, len(input))
  var idx int
  for k,v := range(input) {
    pairs[idx] = &KeyIntValueFloatPair2{k, v}
    idx++
  }
  sort.Sort(&pairs)
  return
}




/* sort by value */
type KeyIntValueIntPair struct {
  Key int
  Value int
}

/* sort by key */
type KeyIntValueIntPair2 struct {
  Key int
  Value int
}

type KeyIntValueIntPairs []*KeyIntValueIntPair
type KeyIntValueIntPairs2 []*KeyIntValueIntPair2

func (pairs *KeyIntValueIntPairs)Swap(i,j int)        {(*pairs)[i], (*pairs)[j] = (*pairs)[j], (*pairs)[i]}
func (pairs *KeyIntValueIntPairs)Less(i,j int) (bool) {return (*pairs)[i].Value > (*pairs)[j].Value}
func (pairs *KeyIntValueIntPairs)Len() (int)          {return len(*pairs)}

func (pairs *KeyIntValueIntPairs2)Swap(i,j int)        {(*pairs)[i], (*pairs)[j] = (*pairs)[j], (*pairs)[i]}
func (pairs *KeyIntValueIntPairs2)Less(i,j int) (bool) {return (*pairs)[i].Key < (*pairs)[j].Key}
func (pairs *KeyIntValueIntPairs2)Len() (int)          {return len(*pairs)}


func SortIntKeys(input map[int]int) (pairs KeyIntValueIntPairs2) {
  pairs = make(KeyIntValueIntPairs2, len(input))
  var idx int
  for k,v := range(input) {
    pairs[idx] = &KeyIntValueIntPair2{k, v}
    idx++
  }
  sort.Sort(&pairs)
  return
}



func SortIntKeysByIntValues(input map[int]int) (pairs KeyIntValueIntPairs) {
  pairs = make(KeyIntValueIntPairs, len(input))
  var idx int
  for k,v := range(input) {
    pairs[idx] = &KeyIntValueIntPair{k, v}
    idx++
  }
  sort.Sort(&pairs)
  return
}

















func GetMaxMapValue(input map[int]int) (key, val int) {
  for k,v := range(input) {
    if v > val {
      val=v
      key=k
    }
  }
  return
}



func SortIntoBins(input map[int][]string, binSize, max int) (bins, binList []int) {
  //initialize bins
  binList = make([]int, int(float64(max)/float64(binSize))+2)
  idx := 0
  for i:= 0 ; idx<len(binList) ; i=i+binSize {
    binList[idx] = i
    idx++
  }
  //sort data
  bins = make([]int, len(binList))
  for freq, instances := range(input) {
    for binIdx, binVal := range(binList) {
      if freq < binVal {
	bins[binIdx-1] += len(instances)
	break
      }
    }
  }
  return
}









/* map1's values are alwasy 1 ; map2's values are probabilities */
func WeightedMapOverlap(map1 map[string]float64, map2 map[string]float64) (overlap float64) {
  for key,_ := range(map1) {
    if _,ok := map2[key] ; ok {
      overlap += map2[key]*map1[key]
    }
  }
  return
}








/* get random key value from map */
func GetRandomKeyValuePair(m map[int]int) (int,int) {
  position := rand.Intn(len(m))
  itr      := 0
  for key,val := range(m) {
    if itr==position {
      return key,val
    }
    itr++
  }
  return -1,-1
}
