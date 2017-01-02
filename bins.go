package d

import (
	"math"
)

func LinearBins(min float64, max float64, n int) D {
	out := D{}
	d := (max - min) / float64(n)
	x := min
	for i := 0; i <= n; i += 1 {
		out = append(out, x+float64(i)*d)
	}
	return out
}

func LogBins(min float64, max float64, n int) D {
	out := D{}
	d := (math.Log10(max) - math.Log10(min)) / float64(n)
	x := math.Log10(min)
	for i := 0; i <= n; i += 1 {
		out = append(out, math.Pow(10, x+float64(i)*d))
	}
	return out
}

func (d D) BinUp(bins D) []int {
	s := d.Sort()
	out := []int{}
	i := 0
	for bIndex := 0; bIndex < len(bins)-1; bIndex++ {
		lower := bins[bIndex]
		upper := bins[bIndex+1]
		out = append(out, 0)
		for ; i < len(s); i += 1 {
			if s[i] >= lower && s[i] < upper {
				out[bIndex] += 1
			} else if s[i] >= upper {
				break
			}
		}
	}
	return out
}

func (d D) OptimalBins(n int) (D, bool) {
	min := d.Min()
	max := d.Max()
	log := false

	if min <= 0 || max <= 0 {
		log = false
	} else if math.Log10(max)-math.Log10(min) > 2 {
		log = true
	}

	if log {
		return LogBins(min, max, n), true
	} else {
		return LinearBins(min, max, n), false
	}
}
