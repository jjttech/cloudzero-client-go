//go:build unit

package costformation

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jjttech/cloudzero-client-go/pkg/config"
)

var _ = Describe("CostFormation", func() {
	Describe("New", func() {
		cfg := config.New()

		c, err := New(cfg)
		Expect(err).To(Succeed())
		Expect(c).NotTo(BeNil())
		Expect(c.client).NotTo(BeNil())
		Expect(c.baseURL).To(Equal(cfg.BaseURL))
	})

	Describe("ReadFile", func() {
		It("errors on an empty filename", func() {
			var c *CostFormation

			d, err := c.ReadFile("")
			Expect(err).To(HaveOccurred())
			Expect(d).To(BeNil())
		})
	})

	Describe("Read", func() {
		It("errors on a nil io.Reader", func() {
			c := &CostFormation{}

			d, err := c.Read(nil)
			Expect(err).To(HaveOccurred())
			Expect(d).To(BeNil())
		})

		d := &Definition{}

		err := d.Read(nil)
		Expect(err).Should(MatchError(ErrInvalidReader))
	})

	Describe("WriteFile", func() {
		It("errors on a nil definition", func() {
			c := &CostFormation{}

			err := c.WriteFile(nil, "")
			Expect(err).Should(MatchError(ErrInvalidDefinition))
		})

		It("succeeds an empty filename (defaults to stdout)", func() {
			c := &CostFormation{}
			d := &Definition{}

			err := c.WriteFile(d, "")
			Expect(err).To(Succeed())
		})
	})

	Describe("Write", func() {
		It("errors on a nil definition", func() {
			c := &CostFormation{}

			err := c.Write(nil, nil)
			Expect(err).Should(MatchError(ErrInvalidDefinition))
		})

		It("errors on a nil io.Reader", func() {
			c := &CostFormation{}
			d := &Definition{}

			err := c.Write(d, nil)
			Expect(err).Should(MatchError(ErrInvalidWriter))
		})
	})
})
