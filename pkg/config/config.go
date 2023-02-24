package config

import (
	"net/http"
	"time"
)

const (
	DefaultBaseURL = "https://api.cloudzero.com/v2"
)

// Config contains information required when talking to the CloudZero API
type Config struct {
	// APIKey used for client authorization
	APIKey string

	// BaseURL used when connecting to CloudZero
	BaseURL string

	// Timeout is the client timeout for HTTP requests.
	Timeout *time.Duration

	// HTTPTransport allows customization of the client's underlying transport.
	HTTPTransport http.RoundTripper

	// UserAgent updates the default user agent string used by the client.
	UserAgent string

	// RetryMax is the maximum attempts to query the API
	RetryMax *int

	// LogLevel can be one of the following values:
	// "panic", "fatal", "error", "warn", "info", "debug", "trace"
	LogLevel string
}

// New returns a Config populated with default parameters and options applied
func New() Config {
	return Config{
		BaseURL:  DefaultBaseURL,
		LogLevel: "info",
	}
}
