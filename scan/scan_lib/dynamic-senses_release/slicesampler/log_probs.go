package slicesampler

import (
  "math"
  "code.google.com/gostat/stat"
)

/* standard normal */
func norm_logpdf(x float64) (y float64) {
  return -math.Pow(x,2)/2.0
}

/* beta */
func beta_logpdf(x float64) (y float64) {
  return stat.Beta_LnPDF(2.0,6.0)(x)
}