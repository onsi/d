package d

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
)

type D []float64

func FromTime(t time.Time) float64 {
	return float64(t.UnixNano())
}

func ToTime(f float64) time.Time {
	return time.Unix(0, int64(f))
}

func FromDuration(d time.Duration) float64 {
	return float64(d)
}

func ToDuration(f float64) time.Duration {
	return time.Duration(f)
}

/* Constructors */

func FromTimes(ts []time.Time) D {
	d := D{}
	for _, t := range ts {
		d = append(d, FromTime(t))
	}
	return d
}

func (d D) ToTimes() []time.Time {
	ts := []time.Time{}
	for _, f := range d {
		ts = append(ts, ToTime(f))
	}
	return ts
}

func FromDurations(ds []time.Duration) D {
	d := D{}
	for _, dur := range ds {
		d = append(d, FromDuration(dur))
	}
	return d
}

func (d D) ToDurations() []time.Duration {
	ds := []time.Duration{}
	for _, f := range d {
		ds = append(ds, ToDuration(f))
	}
	return ds
}

func FromInts(is []int) D {
	d := D{}
	for _, i := range is {
		d = append(d, float64(i))
	}
	return d
}

func (d D) ToInts() []int {
	is := []int{}
	for _, f := range d {
		is = append(is, int(f))
	}
	return is
}

/* Duplicate */

func (d D) Dup() D {
	out := D{}

	for _, f := range d {
		out = append(out, f)
	}

	return out
}

/* Stats */

func (d D) Count() int {
	return len(d)
}

func (d D) Min() float64 {
	if len(d) == 0 {
		return 0
	}

	m := math.MaxFloat64
	for _, f := range d {
		if f < m {
			m = f
		}
	}

	return m
}

func (d D) Max() float64 {
	if len(d) == 0 {
		return 0
	}

	m := -math.MaxFloat64
	for _, f := range d {
		if f > m {
			m = f
		}
	}

	return m
}

func (d D) Sum() float64 {
	s := 0.0

	for _, f := range d {
		s += f
	}

	return s
}

func (d D) Mean() float64 {
	return d.Sum() / float64(d.Count())
}

func (d D) Median() float64 {
	return d.Percentile(0.5)
}

func (d D) Percentile(p float64) float64 {
	if p < 0 || p >= 1.0 {
		panic(fmt.Sprintf("Percentile should be >= 0 and < 1.0, got %f", p))
	}
	sorted := d.Sort()
	index := int(float64(d.Count()) * p)
	return sorted[index]
}

func (d D) Report() string {
	out := ""
	out += fmt.Sprintf("Count: %d\n", d.Count())
	out += fmt.Sprintf("Min: %f\n", d.Min())
	out += fmt.Sprintf("Max: %f\n", d.Max())
	out += fmt.Sprintf("Mean: %f\n", d.Mean())
	out += fmt.Sprintf("Median: %f\n\n", d.Median())

	numBins := d.Count() / 10
	if numBins > 20 {
		numBins = 20
	}
	bins, _ := d.OptimalBins(numBins)
	histogram := d.BinUp(bins)
	max := FromInts(histogram).Max()
	for i := range histogram {

		out += fmt.Sprintf("[%.3f, %.3f)\n%s (%d)\n", bins[i], bins[i+1], strings.Repeat("#", int(float64(20*histogram[i])/max)), histogram[i])
	}

	return out
}

/* Iterators and Filters */

func (d D) Any(cb func(float64) bool) bool {
	for _, f := range d {
		if cb(f) {
			return true
		}
	}
	return false
}

func (d D) None(cb func(float64) bool) bool {
	for _, f := range d {
		if cb(f) {
			return false
		}
	}
	return true
}

func (d D) All(cb func(float64) bool) bool {
	for _, f := range d {
		if !cb(f) {
			return false
		}
	}
	return true
}

func (d D) Filter(cb func(float64) bool) D {
	out := D{}

	for _, f := range d {
		if cb(f) {
			out = append(out, f)
		}
	}

	return out
}

func (d D) Reject(cb func(float64) bool) D {
	out := D{}

	for _, f := range d {
		if !cb(f) {
			out = append(out, f)
		}
	}

	return out
}

func (d D) Each(cb func(float64)) {
	for _, f := range d {
		cb(f)
	}
}

func (d D) Map(cb func(float64) float64) D {
	out := D{}

	for _, f := range d {
		out = append(out, cb(f))
	}

	return out
}

func (d D) Reduce(cb func(memo float64, f float64) float64, s float64) float64 {
	for _, f := range d {
		s = cb(s, f)
	}
	return s
}

// bool operations

func (d D) TruFa(cb func(f float64) bool) []bool {
	out := []bool{}
	for _, f := range d {
		out = append(out, cb(f))
	}
	return out
}

func (d D) ApplyTruFa(trufa []bool) D {
	if len(trufa) != len(d) {
		panic("TruFa nd D must have same length.")
	}

	out := D{}

	for i, b := range trufa {
		if b {
			out = append(out, d[i])
		}
	}

	return out
}

// indexing

func (d D) Index(indices []int) D {
	out := D{}
	for _, i := range indices {
		out = append(out, d[i])
	}

	return out
}

/* Order helpers */

func (d D) Reverse() D {
	out := d.Dup()
	for i, f := range d {
		out[len(d)-1-i] = f
	}
	return out
}

func (d D) Sort() D {
	dup := d.Dup()
	sort.Float64s(dup)
	return dup
}

type indexSorter struct {
	indices []int
	d       D
}

func (s indexSorter) Len() int { return len(s.d) }
func (s indexSorter) Less(i, j int) bool {
	return s.d[i] < s.d[j] || math.IsNaN(s.d[i]) && !math.IsNaN(s.d[j])
}
func (s indexSorter) Swap(i, j int) {
	s.d[i], s.d[j] = s.d[j], s.d[i]
	s.indices[i], s.indices[j] = s.indices[j], s.indices[i]
}

func (d D) SortedIndices() []int {
	i := indexSorter{
		Range(d.Count()),
		d.Dup(),
	}
	sort.Sort(i)
	return i.indices
}
