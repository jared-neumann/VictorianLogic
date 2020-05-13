package slicesampler

import "math"
import "math/rand"


func Sample(old_val float64, log_prob func(float64)float64, step_size, max_step, lower, upper float64, log_prob_old_val float64) (new_val float64) {
  /* check validity of arguments */
  
  /* find log density at initial point if not known */
  if (log_prob_old_val == 0.0) {
    log_prob_old_val = log_prob(old_val)
  }
  
  /* determine the slice level = [[log_prob_old_val - random_number]]  */
  logy := log_prob_old_val - rand.ExpFloat64()
  println(old_val, log_prob_old_val, logy)

  /* find the original interval to sample from */
  u := rand.Float64() * step_size
  left  := old_val - u
  right := old_val + (step_size-u)
  
  println("slice level done")
  /* expand interval until ends are outside slice, or step limit  is reached*/
  if math.IsInf(max_step,0) {
    for ; ; {
      println("L", left, log_prob(left), logy, lower)
      if left <=lower {break}
      if log_prob(left) <= logy {break}
      left = left-step_size
    }
    for ; ; {
      println("R", "(",right,")", log_prob(right), ">=", logy, "or", right, ">=", upper, step_size)
      if right >= upper {break}
      if log_prob(right) <= logy {break}
      right = right+step_size
    }
    
  } else if max_step>1 {
    j := math.Floor(rand.Float64()*max_step)
    k := (max_step-1)-j
    
    for ; j>0 ; j-- {
      if left<=lower {break}
      if log_prob(left) <= logy {break}
      left = left-step_size
    }
    for ; k>0 ; k-- {
      if right>=upper {break}
      if log_prob(right) >=logy {break}
      right = right+step_size
    }
  }
  println("expansion done")
  /* shrink the interval to lower and upper bounds */
  if left<lower{
    left=lower
  }
  if right>upper{
    right=upper
  }
  println("shrink done")
  /* sample from the interval, shrinking it on each rejection */
  for ; ; {
    new_val = rand.Float64() * (right-left) + left
    log_prob_new_val := log_prob(new_val)
    println(left,right,new_val,log_prob_new_val,logy)
    if log_prob_new_val>=logy {break}
    
    if new_val>old_val {
      right = new_val
    } else {
      left = new_val
    }
  }
  println("sampling done")
  return new_val
}
