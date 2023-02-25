package cloudzero

import (
	"errors"
	"net/http"
	"strings"
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
func WithHTTPTimeout(t time.Duration) ConfigOption {
	return func(cfg *config.Config) error {
		if t < 0 {
			return errors.New("invalid HTTP Timeout")
		}

		var timeout = &t
		cfg.Timeout = timeout
		return nil
	}
}

// WithHTTPTransport sets the HTTP Transporter.
func WithHTTPTransport(transport http.RoundTripper) ConfigOption {
	return func(cfg *config.Config) error {
		if nil == transport {
			return errors.New("HTTP Transport can not be nil")
		}

		cfg.HTTPTransport = transport
		return nil
	}
}

// WithUserAgent sets the HTTP UserAgent for API requests.
func WithUserAgent(ua string) ConfigOption {
	return func(cfg *config.Config) error {
		if "" == ua {
			return errors.New("user-agent can not be empty")
		}

		cfg.UserAgent = ua
		return nil
	}
}

// WithBaseURL sets the base URL used to make requests to the REST API V2.
func WithBaseURL(url string) ConfigOption {
	return func(cfg *config.Config) error {
		if "" == url {
			return errors.New("base URL can not be empty")
		}

		cfg.BaseURL = url
		return nil
	}
}

// WithLogLevel sets the log level for the client.
func WithLogLevel(logLevel string) ConfigOption {
	return func(cfg *config.Config) error {
		llv := strings.ToLower(logLevel)

		switch llv {
		case "panic", "fatal", "error", "warn", "warning", "info", "debug", "trace":
			cfg.LogLevel = llv
		default:
			return errors.New("invalid log level")
		}

		return nil
	}
}
