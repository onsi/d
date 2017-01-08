package d_test

import (
	. "github.com/onsi/d"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "D Suite")
}

var entries []SchemaEntry
var schema Schema

var _ = BeforeEach(func() {
	entries = []SchemaEntry{
		{"foo", FloatType},
		{"bar", StringType},
		{"baz", FloatType},
	}

	schema = NewSchema(entries)
})
