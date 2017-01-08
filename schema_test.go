package d_test

import (
	. "github.com/onsi/d"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Schema", func() {
	Describe("Creating Schemas", func() {
		It("creates the schema ", func() {
			Ω(schema.Entries()).Should(Equal(entries))
		})
	})

	Describe("Validate", func() {
		Context("when all is good", func() {
			It("succeeds", func() {
				Ω(schema.Validate([]Value{
					VF(1.0), VS("abc"), VF(-3.0),
				})).Should(Succeed())
			})
		})

		Context("when the length is wrong", func() {
			It("errors", func() {
				Ω(schema.Validate([]Value{
					VF(1.0), VS("abc"),
				})).Should(MatchError("length mismatch"))
			})
		})

		Context("when the types are wrong", func() {
			It("errors", func() {
				Ω(schema.Validate([]Value{
					VF(1.0), VS("abc"), VS("abc"),
				})).Should(MatchError("type mismatch"))
			})
		})
	})

	Describe("Keys", func() {
		It("returns the list of keys", func() {
			Ω(schema.Keys()).Should(Equal(S{"foo", "bar", "baz"}))
		})
	})

	Describe("Length", func() {
		It("returns the number of keys", func() {
			Ω(schema.Length()).Should(Equal(3))
		})
	})

	Describe("Index", func() {
		It("return the index for the given key", func() {
			Ω(schema.Index("foo")).Should(Equal(0))
			Ω(schema.Index("bar")).Should(Equal(1))
			Ω(schema.Index("baz")).Should(Equal(2))
		})

		Context("when the key cannot be found", func() {
			It("returns NOT_FOUND", func() {
				Ω(schema.Index("wibble")).Should(Equal(NOT_FOUND))
			})
		})
	})

	Describe("Type", func() {
		It("returns the type for the given Key", func() {
			Ω(schema.Type("foo")).Should(Equal(FloatType))
			Ω(schema.Type("bar")).Should(Equal(StringType))
			Ω(schema.Type("baz")).Should(Equal(FloatType))
		})

		Context("when the key cannot be found", func() {
			It("returns the NoneType", func() {
				Ω(schema.Type("wibble")).Should(Equal(NoneType))
			})
		})
	})
})
