//go:build unit

package cloudzero

import (
	"net/http"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jjttech/cloudzero-client-go/pkg/config"
)

var _ = Describe("Option", func() {
	Describe("WithAPIKey", func() {
		It("sets the API key", func() {
			cfg := config.Config{}

			fn := WithAPIKey("test-api-key")
			err := fn(&cfg)

			Expect(err).To(Succeed())
			Expect(cfg.APIKey).To(Equal("test-api-key"))
		})
	})

	Describe("WithHTTPTimeout", func() {
		It("sets the HTTP timeout", func() {
			cfg := config.Config{}

			fn := WithHTTPTimeout(1 * time.Minute)
			err := fn(&cfg)

			Expect(err).To(Succeed())
			Expect(cfg.Timeout).NotTo(BeNil())
			Expect(*cfg.Timeout).To(Equal(1 * time.Minute))
		})

		It("errors on timeouts less than zero", func() {
			cfg := config.Config{}

			fn := WithHTTPTimeout(-1 * time.Minute)
			err := fn(&cfg)

			Expect(err).To(HaveOccurred())
			Expect(cfg.Timeout).To(BeNil())
		})
	})

	Describe("WithHTTPTransport", func() {
		It("errors on nil HTTP Transport", func() {
			cfg := config.Config{}

			fn := WithHTTPTransport(nil)
			err := fn(&cfg)

			Expect(err).To(HaveOccurred())
		})

		It("sets the HTTP Transport", func() {
			cfg := config.Config{}

			fn := WithHTTPTransport(http.DefaultTransport)
			err := fn(&cfg)

			Expect(err).To(Succeed())
			Expect(cfg.HTTPTransport).To(Equal(http.DefaultTransport))
		})
	})

	Describe("WithUserAgent", func() {
		It("errors on empty User Agent", func() {
			cfg := config.Config{}

			fn := WithUserAgent("")
			err := fn(&cfg)

			Expect(err).To(HaveOccurred())
		})

		It("sets the User Agent", func() {
			cfg := config.Config{}

			fn := WithUserAgent("user-agent")
			err := fn(&cfg)

			Expect(err).To(Succeed())
			Expect(cfg.UserAgent).To(Equal("user-agent"))
		})
	})

	Describe("WithBaseURL", func() {
		It("erros on empty Base URL", func() {
			cfg := config.Config{}

			fn := WithBaseURL("")
			err := fn(&cfg)

			Expect(err).To(HaveOccurred())
		})

		It("sets the Base URL", func() {
			cfg := config.Config{}

			fn := WithBaseURL("http://localhost/")
			err := fn(&cfg)

			Expect(err).To(Succeed())
			Expect(cfg.BaseURL).To(Equal("http://localhost/"))
		})
	})

	Describe("WithLogLevel", func() {
		DescribeTable("sets the log level",
			func(level string, isErr bool) {
				cfg := config.Config{}
				fn := WithLogLevel(level)
				err := fn(&cfg)

				if isErr {
					Expect(err).To(HaveOccurred())
				} else {
					Expect(err).To(Succeed())
					Expect(cfg.LogLevel).To(Equal(strings.ToLower(level)))
				}
			},
			Entry("Empty String", "", true),
			Entry("Not a level", "bowtie", true),
			Entry("Not a level", "invalid", true),
			Entry("Level: panic", "panic", false),
			Entry("Level: PANIC", "PANIC", false),
			Entry("Level: fatal", "fatal", false),
			Entry("Level: FATAL", "FATAL", false),
			Entry("Level: error", "error", false),
			Entry("Level: ERROR", "ERROR", false),
			Entry("Level: warn", "warn", false),
			Entry("Level: WARN", "WARN", false),
			Entry("Level: warning", "warning", false),
			Entry("Level: WARNING", "WARNING", false),
			Entry("Level: info", "info", false),
			Entry("Level: INFO", "INFO", false),
			Entry("Level: debug", "debug", false),
			Entry("Level: DEBUG", "DEBUG", false),
			Entry("Level: trace", "trace", false),
			Entry("Level: TRACE", "TRACE", false),
		)
	})
})
