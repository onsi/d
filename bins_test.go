package d_test

import (
	. "github.com/onsi/d"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bins", func() {
	Describe("LinearBins", func() {
		It("generates linear bins", func() {
			Ω(LinearBins(-2, 10, 6)).Should(Equal(D{
				-2, 0, 2, 4, 6, 8, 10,
			}))
		})
	})

	Describe("LogBins", func() {
		It("generates log bins", func() {
			Ω(LogBins(1e-3, 1e2, 5)).Should(Equal(D{
				1e-3, 1e-2, 1e-1, 1, 10, 100,
			}))
		})
	})

	Describe("D.BinUp", func() {
		It("bins up the data set by passed in bins", func() {
			dataSet := D{
				-17, -1, -0.5, 1, 3, 4, 5, 8, 8.5, 9, 10, 17,
			}
			dataSet = dataSet.Index(Permutation(dataSet.Count()))

			Ω(dataSet.BinUp(LinearBins(-2, 10, 6))).Should(Equal([]int{
				2, //[-2,0)
				1, //[0,2)
				1, //[2,4)
				2, //[4,6)
				0, //[6,8)
				3, //[8,10)
			}))

			dataSet = D{
				0, 1,
			}

			Ω(dataSet.BinUp(LinearBins(-2, 10, 6))).Should(Equal([]int{
				0, //[-2,0)
				2, //[0,2)
				0, //[2,4)
				0, //[4,6)
				0, //[6,8)
				0, //[8,10)
			}))

		})
	})

	Describe("D.OptimalBins", func() {
		It("returns the optimal set of bins", func() {
			bins, log := D{
				-1e3,
				1,
			}.OptimalBins(10)
			Ω(bins).Should(Equal(LinearBins(-1e3, 1, 10)))
			Ω(log).Should(BeFalse())

			bins, log = D{
				-1,
				10,
			}.OptimalBins(10)
			Ω(bins).Should(Equal(LinearBins(-1, 10, 10)))
			Ω(log).Should(BeFalse())

			bins, log = D{
				1,
				100,
			}.OptimalBins(10)
			Ω(bins).Should(Equal(LinearBins(1, 100, 10)))
			Ω(log).Should(BeFalse())

			bins, log = D{
				1,
				101,
			}.OptimalBins(10)
			Ω(bins).Should(Equal(LogBins(1, 101, 10)))
			Ω(log).Should(BeTrue())
		})
	})
})
