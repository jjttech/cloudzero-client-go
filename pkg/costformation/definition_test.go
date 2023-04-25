//go:build unit

package costformation

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Definition", func() {
	Describe("ReadFile", func() {
		It("errors on a nil definition", func() {
			var d *Definition

			err := d.ReadFile("")
			Expect(err).Should(MatchError(ErrInvalidDefinition))
		})
	})

	Describe("Read", func() {
		It("errors on a nil definition", func() {
			var d *Definition

			err := d.Read(nil)
			Expect(err).Should(MatchError(ErrInvalidDefinition))
		})

		It("errors on a nil io.Reader", func() {
			d := &Definition{}

			err := d.Read(nil)
			Expect(err).Should(MatchError(ErrInvalidReader))
		})
	})

	Describe("WriteFile", func() {
		It("errors on a nil definition", func() {
			var d *Definition

			err := d.WriteFile("")
			Expect(err).Should(MatchError(ErrInvalidDefinition))
		})
	})

	Describe("Write", func() {
		It("errors on a nil definition", func() {
			var d *Definition

			err := d.Write(nil)
			Expect(err).Should(MatchError(ErrInvalidDefinition))
		})

		It("errors on a nil io.Reader", func() {
			d := &Definition{}

			err := d.Write(nil)
			Expect(err).Should(MatchError(ErrInvalidWriter))
		})
	})
})
