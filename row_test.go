package d_test

import (
	. "github.com/onsi/d"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Row", func() {
	var row Row

	BeforeEach(func() {
		row = NewRow(schema, []Value{
			VF(1.0),
			VS("abc"),
			VF(-3.0),
		})
	})

	Describe("Creation", func() {
		Context("when validation fails", func() {
			It("panics", func() {
				Ω(func() {
					NewRow(schema, []Value{VF(1.0)})
				}).Should(Panic())
			})
		})
	})

	Describe("Schema", func() {
		It("returns the schema", func() {
			Ω(row.Schema()).Should(Equal(schema))
		})
	})

	Describe("Keys", func() {
		It("returns the keys", func() {
			Ω(row.Keys()).Should(Equal(S{"foo", "bar", "baz"}))
		})
	})

	Describe("Length", func() {
		It("returns the length of the row", func() {
			Ω(row.Length()).Should(Equal(3))
		})
	})

	Describe("Getting Entries", func() {
		Describe("Value", func() {
			It("gets the value for the given key", func() {
				Ω(row.Value("foo")).Should(Equal(VF(1.0)))
				Ω(row.Value("bar")).Should(Equal(VS("abc")))
				Ω(row.Value("baz")).Should(Equal(VF(-3.0)))
			})

			Context("when the value does not exist", func() {
				It("panics", func() {
					Ω(func() {
						row.Value("wibble")
					}).Should(Panic())
				})
			})
		})
	})

	Describe("Setting Entries", func() {
		Describe("SetValue", func() {
			It("should set the value", func() {
				Ω(row.SetValue("foo", VF(17.0)).Value("foo")).Should(Equal(VF(17.0)))
			})

			Context("when the key does not exist", func() {
				It("panics", func() {
					Ω(func() {
						row.SetValue("wibble", VF(17.0))
					}).Should(Panic())
				})
			})

			Context("when the value has the wrong type", func() {
				It("panics", func() {
					Ω(func() {
						row.SetValue("foo", VS("abc"))
					}).Should(Panic())
				})
			})
		})
	})
})
