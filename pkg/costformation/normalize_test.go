//go:build unit

package costformation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Normalize", func() {
	It("should not change a basic string", func() {
		x := "asdf"
		y := Normalize(x)
		Expect(y).To(Equal(x))
	})

	It("should lowercase a string", func() {
		x := "AsDF"
		y := Normalize(x)
		Expect(y).To(Equal("asdf"))
	})

	It("should convert special characters to a dash", func() {
		chars := `.,/#!$%^&*;:=_~()\'`
		for _, v := range chars {
			y := Normalize(string(v))
			Expect(y).To(Equal("-"))
		}
	})

	// From https://docs.cloudzero.com/docs/cfdl-reference#normalize
	It("should conform to the example in the CloudZero docs", func() {
		y := Normalize("Production/Resources#4561")
		Expect(y).To(Equal("production-resources-4561"))
	})
})
