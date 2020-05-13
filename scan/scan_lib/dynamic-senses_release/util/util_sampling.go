package util

import (
       "math/rand"
       "fmt"
       "math"
)

/* Function for obtaining a cumulative sample    *
 * from an unnormalized probability distribution */
func GetSample(dist []float64) (idx int) {
    if len(dist) <= 0 {
        panic("sample from empty distribution")
    }
    sum := 0.0
    for _, v := range dist {
        if v < 0 {
            panic(fmt.Sprintf("bad dist: %v", dist))
        }
        sum += v
    }
    u := rand.Float64() * sum
    sum = 0
    for i, v := range dist {
        sum += v
        if u < sum {
            return i
        }
    }
    fmt.Println(dist)
    panic("sampleDiscrete gets out of all possiblilities")
}


/* Function for obtaining a cumulative sample
 *  from a CUMULATIVE PROBABILITY DENSITY */
func GetCumulativeSample(cumul []float64) (int) {
  //sample random number in [0:1] and normalize dist
  i := rand.Float64()
  s := i * cumul[len(cumul)-1]
  //find corresponding topic
  for top,_ := range(cumul) {
    if cumul[top] > s {
      return top
    }
  }
  return -1
}


/* Function for obtaining a cumulative sample
 *  from an unnormalized probability distribution */
func GetSampleLog(mult []float64) (idx int) {
  _,b := FindMax(mult)
  z := expSum(mult,b)
  //sample random number in [0:1] and normalize dist
  i := rand.Float64()
  //find corresponding topic
  sum := 0.0
  for top,val := range(mult) {
    norm := ExpNormalize(val, b, z)
    sum += norm
    if sum > i {
      return top
    }
  }
  return -1
}

/* Function for obtaining a cumulative sample
 *  from a normalized probability distribution */
func GetSampleLog_normalized(mult []float64) (idx int) {
  //sample random number in [0:1] and normalize dist
  i := rand.Float64()
  //find corresponding topic
//   fmt.Println(i)
  sum := 0.0
  for top,val := range(mult) {
    sum += math.Exp(val)
    if sum > i {
      return top
    }
  }
  return -1
}


/* Function for obtaining a cumulative sample
 *  from an unnormalized probability distribution */
func GetSampleLog_gen(generator *rand.Rand, mult []float64) (idx int) {
  _,b := FindMax(mult)
  z := expSum(mult,b)
  //sample random number in [0:1] and normalize dist
  i := generator.Float64()
  //find corresponding topic
  sum := 0.0
  for top,val := range(mult) {
    norm := ExpNormalize(val, b, z)
    sum += norm
    if sum > i {
      return top
    }
  }
  return -1
}


/* Function for obtaining NUM samples
 *  from an CUMULATIVE PROBABILITY DENSITY 
 *  (1) sort random numbers 
 *  (2) iterate only once over cumulative fct
 *  sounds nice and is what ppl do in the literature but sorting takes too long */
func GetRandomSamples(cumul []float64, n int) (samples []int) {
  //sample random number in [0:1] and normalize dist
  samples = make([]int    , len(cumul))
  rands := GenerateOrderedRandomNumbers(n)
  //find corresponding indices for random numbers
  k:=0
  for i:=0 ; i<n ; i++ {
    for  ; cumul[k]<rands[i] ;  {
      k++
    }
    samples[k]++
  }
  return
}


/* Function for obtaining NUM samples
 *  from an CUMULATIVE PROBABILITY DENSITY 
 *  (1) sort random numbers 
 *  (2) iterate only once over cumulative fct
 *  sounds nice and is what ppl do in the literature but sorting takes too long */
func GetRandomSamples2(cumul []float64, n int) (samples map[int]int) {
  //sample random number in [0:1] and normalize dist
  samples = make(map[int]int)
  rands := GenerateOrderedRandomNumbers(n)
  //find corresponding indices for random numbers
  k:=0
  for i:=0 ; i<n ; i++ {
    for  ; cumul[k]<rands[i] ;  {
      k++
    }
    samples[k]++
  }
  return
}

func GetStratifiedSamples(cumul []float64, n int) (samples []int) {
  // init
  samples = make([]int    , len(cumul))
  interval := 1.0 / float64(n)
  k  := 0
  rr := 0.0
  // for each interval, sample random number and find corresponding element (index) in cumul
  for i:=0 ; i<n ; i++ {
    // sample random number in the right interval
    rr = ScaleProbability(rand.Float64(), [2]float64{float64(i)*interval, (float64(i+1))*interval})
    for  ; cumul[k]<rr ;  {
        k++
    }
      samples[k]++
  }
  return
}


func GetSystematicSamples(cumul []float64, n int) (samples []int) {
  // init
  samples = make([]int    , len(cumul))
  interval := 1.0 / float64(n)
  k  := 0
  rr := rand.Float64()
  // for each interval, sample random number and find corresponding element (index) in cumul
  for i:=0 ; i<n ; i++ {
    // sample random number in the right interval
    ri := ScaleProbability(rr, [2]float64{float64(i)*interval, (float64(i+1))*interval})
    for  ; cumul[k]<ri ;  {
        k++
    }
      samples[k]++
  }
  return
}




func GetSystematicSamples_log(log_dist []float64, n int) (samples []int) {
  /* logsumexp and cumulate */
  dist := Exp_copy(log_dist)
  cumul:= Cumulate_copy(dist)
  // init
  samples = make([]int    , len(cumul))
  interval := 1.0 / float64(n)
  k  := 0
  rr := rand.Float64()
  // for each interval, sample random number and find corresponding element (index) in cumul
  for i:=0 ; i<n ; i++ {
    // sample random number in the right interval
    ri := ScaleProbability(rr, [2]float64{float64(i)*interval, (float64(i+1))*interval})
    for  ; cumul[k]<ri ;  {
        k++
    }
      samples[k]++
  }
  return
}
