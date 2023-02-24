package cloudzero

import (
	"errors"
	"net/http"
	"time"

	"github.com/jjttech/cloudzero-client-go/pkg/config"
)

// ConfigOption allows for setting config options
type ConfigOption func(*config.Config) error

func WithAPIKey(apiKey string) ConfigOption {
	return func(cfg *config.Config) error {
		cfg.APIKey = apiKey
		return nil
	}
}

// WithHTTPTimeout sets the timeout for HTTP requests.
func ConfigHTTPTimeout(t time.Duration) ConfigOption {
	return func(cfg *config.Config) error {
		var timeout = &t
		cfg.Timeout = timeout
		return nil
	}
}

// WithHTTPTransport sets the HTTP Transporter.
func WithHTTPTransport(transport http.RoundTripper) ConfigOption {
	return func(cfg *config.Config) error {
		if transport != nil {
			cfg.HTTPTransport = transport
			return nil
		}

		return errors.New("HTTP Transport can not be nil")
	}
}

// WithUserAgent sets the HTTP UserAgent for API requests.
func WithUserAgent(ua string) ConfigOption {
	return func(cfg *config.Config) error {
		if ua != "" {
			cfg.UserAgent = ua
			return nil
		}

		return errors.New("user-agent can not be empty")
	}
}

// WithBaseURL sets the base URL used to make requests to the REST API V2.
func WithBaseURL(url string) ConfigOption {
	return func(cfg *config.Config) error {
		if url != "" {
			cfg.BaseURL = url
			return nil
		}

		return errors.New("base URL can not be empty")
	}
}

// WithLogLevel sets the log level for the client.
func WithLogLevel(logLevel string) ConfigOption {
	return func(cfg *config.Config) error {
		if logLevel != "" {
			cfg.LogLevel = logLevel
			return nil
		}

		return errors.New("log level can not be empty")
	}
}
