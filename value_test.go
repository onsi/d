package d_test

import (
	. "github.com/onsi/d"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Value", func() {
	Describe("VF", func() {
		It("generates float values", func() {
			Ω(VF(1)).Should(Equal(Value{
				F:    1,
				Type: FloatType,
			}))

			Ω(VF(1, true)).Should(Equal(Value{
				Type:  FloatType,
				Blank: true,
			}))
		})
	})

	Describe("VS", func() {
		It("generates string values", func() {
			Ω(VS("foo")).Should(Equal(Value{
				S:    "foo",
				Type: StringType,
			}))

			Ω(VS("foo", true)).Should(Equal(Value{
				Type:  StringType,
				Blank: true,
			}))
		})
	})
})
