//go:build unit

package cloudzero

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CloudZero", func() {
	Describe("New", func() {
		It("returns a new client without any options", func() {
			cz, err := New()

			Expect(err).To(Succeed())
			Expect(cz).NotTo(BeNil())
		})

		It("returns a new client with an option set", func() {
			cz, err := New(WithAPIKey("test-api-key"))

			Expect(err).To(Succeed())
			Expect(cz).NotTo(BeNil())
		})

		It("returns an error on an invalid option", func() {
			cz, err := New(WithHTTPTimeout(-1 * time.Minute))

			Expect(err).To(HaveOccurred())
			Expect(cz).To(BeNil())
		})
	})
})
