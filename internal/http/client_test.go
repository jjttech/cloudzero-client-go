//go:build unit

package http

import (
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jjttech/cloudzero-client-go/pkg/config"
)

var _ = Describe("Client", func() {
	Describe("NewClient", func() {
		It("returns a new client without any options", func() {
			cfg := config.New()

			client, err := NewClient(cfg)

			Expect(err).To(Succeed())
			Expect(client).NotTo(BeNil())
			Expect(client.client).NotTo(BeNil())
			Expect(client.client.HTTPClient).NotTo(BeNil())

			// Defaults set by NewClient
			Expect(client.client.RetryMax).To(Equal(defaultRetryMax))                   // default retries
			Expect(client.client.Logger).To(BeNil())                                    // Lower logger disabled
			Expect(client.client.HTTPClient.Timeout).To(Equal(defaultTimeout))          // default timeout
			Expect(client.client.HTTPClient.Transport).To(Equal(http.DefaultTransport)) // default transport config
		})

		It("copies the Timeout from config", func() {
			cfg := config.New()
			t := 1 * time.Minute
			cfg.Timeout = &t

			client, err := NewClient(cfg)

			Expect(err).To(Succeed())
			Expect(client).NotTo(BeNil())
			Expect(client.client).NotTo(BeNil())
			Expect(client.client.HTTPClient).NotTo(BeNil())

			Expect(client.client.HTTPClient.Timeout).To(Equal(*cfg.Timeout))
		})

		It("copies the Transport from config", func() {
			cfg := config.New()
			t := http.Transport{}
			cfg.HTTPTransport = &t

			client, err := NewClient(cfg)

			Expect(err).To(Succeed())
			Expect(client).NotTo(BeNil())
			Expect(client.client).NotTo(BeNil())
			Expect(client.client.HTTPClient).NotTo(BeNil())

			Expect(client.client.HTTPClient.Transport).To(Equal(&t))
		})

		It("copies RetryMax from config", func() {
			cfg := config.New()
			t := 5
			cfg.RetryMax = &t

			client, err := NewClient(cfg)

			Expect(err).To(Succeed())
			Expect(client).NotTo(BeNil())
			Expect(client.client).NotTo(BeNil())

			Expect(client.client.RetryMax).To(Equal(t))
		})

		It("copies User Agent from config", func() {
			cfg := config.New()
			cfg.UserAgent = "test-user-agent"

			client, err := NewClient(cfg)

			Expect(err).To(Succeed())
			Expect(client).NotTo(BeNil())

			Expect(client.userAgent).To(Equal("test-user-agent"))
		})

		It("copies API Key from config", func() {
			cfg := config.New()
			cfg.APIKey = "test-api-key"

			client, err := NewClient(cfg)

			Expect(err).To(Succeed())
			Expect(client).NotTo(BeNil())

			Expect(client.apiKey).To(Equal("test-api-key"))
		})
	})
})
