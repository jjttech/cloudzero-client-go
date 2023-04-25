//go:build unit

package config

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	Describe("New", func() {
		It("creates a new Config", func() {
			cfg := New()

			Expect(cfg.BaseURL).To(Equal(DefaultBaseURL))
			Expect(cfg.LogLevel).To(Equal(DefaultLogLevel))
		})
	})
})
