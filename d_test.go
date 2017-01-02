package d_test

import (
	. "github.com/onsi/d"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("D", func() {
	var _ = It("tests a bunch of D-related things", func() {
		empty := D{}
		d := D{0.2, -1.3, 12.4, -3.9, 9.8}

		Ω(empty.Count()).Should(Equal(0))
		Ω(d.Count()).Should(Equal(len(d)))

		Ω(empty.Min()).Should(Equal(0.0))
		Ω(empty.Max()).Should(Equal(0.0))

		Ω(d.Min()).Should(Equal(-3.9))
		Ω(d.Max()).Should(Equal(12.4))

		Ω(d.Sum()).Should(Equal(d[0] + d[1] + d[2] + d[3] + d[4]))

		Ω(d.Sort()).Should(Equal(D{-3.9, -1.3, 0.2, 9.8, 12.4}))
		Ω(d[0]).Should(Equal(0.2), "Sort should not mutate d")

		Ω(d.Median()).Should(Equal(0.2))
		Ω(d.Percentile(0)).Should(Equal(-3.9))
		Ω(d.Percentile(0.6)).Should(Equal(9.8))
		Ω(d.Percentile(0.8)).Should(Equal(12.4))
		Ω(func() { d.Percentile(-0.1) }).Should(Panic())
		Ω(func() { d.Percentile(1.0) }).Should(Panic())

		Ω(d.Any(func(f float64) bool {
			return f > 12.4
		})).Should(BeFalse())

		Ω(d.Any(func(f float64) bool {
			return f == 12.4
		})).Should(BeTrue())

		Ω(d.None(func(f float64) bool {
			return f > 12.4
		})).Should(BeTrue())

		Ω(d.None(func(f float64) bool {
			return f == 12.4
		})).Should(BeFalse())

		Ω(d.All(func(f float64) bool {
			return f <= 12.4
		})).Should(BeTrue())

		Ω(d.All(func(f float64) bool {
			return f <= 11.4
		})).Should(BeFalse())

		Ω(d.Filter(func(f float64) bool {
			return f < 12.4 && f >= 0
		})).Should(Equal(D{0.2, 9.8}))

		Ω(d.Reject(func(f float64) bool {
			return f < 12.4 && f >= 0
		})).Should(Equal(D{-1.3, 12.4, -3.9}))

		s := 0.0
		d.Each(func(f float64) {
			s += f
		})
		Ω(s).Should(Equal(d.Sum()))

		out := d.Map(func(f float64) float64 {
			return -f
		})
		Ω(out).Should(Equal(D{-0.2, 1.3, -12.4, 3.9, -9.8}))

		trufa := d.TruFa(func(f float64) bool {
			return f < 12.4 && f >= 0
		})

		Ω(trufa).Should(Equal([]bool{true, false, false, false, true}))
		Ω(d.ApplyTruFa(trufa)).Should(Equal(D{0.2, 9.8}))

		s = d.Reduce(func(memo float64, f float64) float64 {
			return memo + f
		}, 0)
		Ω(s).Should(Equal(d.Sum()))

		Ω(d.Reverse()).Should(Equal(D{9.8, -3.9, 12.4, -1.3, 0.2}))

		Ω(d.Index([]int{
			0, 0, 1, 3, 1,
		})).Should(Equal(D{0.2, 0.2, -1.3, -3.9, -1.3}))

		Ω(d.Index(d.SortedIndices())).Should(Equal(d.Sort()))

		indices := Range(d.Count())
		reveredIndices := FromInts(indices).Reverse().ToInts()
		Ω(d.Index(reveredIndices)).Should(Equal(d.Reverse()))
	})
})
