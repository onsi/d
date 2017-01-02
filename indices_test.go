package d_test

import (
	. "github.com/onsi/d"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Indices", func() {
	It("tests a bunch of indices related things", func() {
		Ω(Range(5)).Should(Equal([]int{0, 1, 2, 3, 4}))

		permuted := Permutation(100)
		Ω(permuted).ShouldNot(Equal(Range(100)))
		Ω(permuted).Should(ConsistOf(Range(100)))
	})
})
