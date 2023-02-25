//go:build unit

package http

import (
	"context"
	"net/http"
	"time"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

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

	Describe("setHeaders", func() {
		It("sets default headers only", func() {
			client := Client{}
			req, err := retryablehttp.NewRequest(http.MethodGet, "http://localhost", nil)
			Expect(err).To(Succeed())

			err = client.setHeaders(req)
			Expect(err).To(Succeed())

			// Unset
			h := req.Header.Get("Authorization")
			Expect(h).To(Equal(""))

			// Default user agent
			h = req.Header.Get("User-Agent")
			Expect(h).To(Equal(defaultUserAgent))

			// JSON
			h = req.Header.Get("Content-Type")
			Expect(h).To(Equal("application/json"))
		})

		It("sets the Authorization header", func() {
			client := Client{
				apiKey: "test-api-key",
			}
			req, err := retryablehttp.NewRequest(http.MethodGet, "http://localhost", nil)
			Expect(err).To(Succeed())

			err = client.setHeaders(req)
			Expect(err).To(Succeed())

			h := req.Header.Get("Authorization")
			Expect(h).To(Equal("test-api-key"))
		})

		It("sets the User-Agent header", func() {
			client := Client{
				userAgent: "test-user-agent",
			}
			req, err := retryablehttp.NewRequest(http.MethodGet, "http://localhost", nil)
			Expect(err).To(Succeed())

			err = client.setHeaders(req)
			Expect(err).To(Succeed())

			h := req.Header.Get("User-Agent")
			Expect(h).To(Equal("test-user-agent"))
		})
	})

	Describe("Client Actions", func() {
		var server *ghttp.Server

		BeforeEach(func() {
			server = ghttp.NewServer()
		})

		AfterEach(func() {
			server.Close()
		})

		Describe("Get", func() {
			BeforeEach(func() {
				server.AppendHandlers(
					ghttp.VerifyRequest("GET", "/"),
				)
			})

			It("should make a basic GET request", func() {
				ctx := context.TODO()
				client, _ := NewClient(config.Config{})

				resp, err := client.Get(ctx, server.URL())
				Expect(err).To(Succeed())
				Expect(resp).NotTo(BeNil())
			})
		})

		Describe("Post", func() {
			BeforeEach(func() {
				server.AppendHandlers(
					ghttp.VerifyRequest("POST", "/"),
				)
			})

			It("should make a basic POST request", func() {
				ctx := context.TODO()
				client, _ := NewClient(config.Config{})

				resp, err := client.Post(ctx, server.URL(), nil)
				Expect(err).To(Succeed())
				Expect(resp).NotTo(BeNil())
			})
		})

		Describe("Put", func() {
			BeforeEach(func() {
				server.AppendHandlers(
					ghttp.VerifyRequest("PUT", "/"),
				)
			})

			It("should make a basic PUT request", func() {
				ctx := context.TODO()
				client, _ := NewClient(config.Config{})

				resp, err := client.Put(ctx, server.URL(), nil)
				Expect(err).To(Succeed())
				Expect(resp).NotTo(BeNil())
			})
		})

		Describe("Delete", func() {
			BeforeEach(func() {
				server.AppendHandlers(
					ghttp.VerifyRequest("DELETE", "/"),
				)
			})

			It("should make a basic DELETE request", func() {
				ctx := context.TODO()
				client, _ := NewClient(config.Config{})

				resp, err := client.Delete(ctx, server.URL())
				Expect(err).To(Succeed())
				Expect(resp).NotTo(BeNil())
			})
		})
	})
})
