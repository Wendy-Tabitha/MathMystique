package mathskills

import "math"

func Variance(num []float64) float64 {
	n := len(num)
	sum := 0.0
	numb := Average(num)
	for _, nb := range num {
		sum += math.Pow(numb-nb, 2)
	}
	return sum / float64(n)
}
