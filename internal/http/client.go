package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	retryablehttp "github.com/hashicorp/go-retryablehttp"

	"github.com/jjttech/cloudzero-client-go/internal/version"
	"github.com/jjttech/cloudzero-client-go/pkg/config"
)

/*
 * CloudZero appears to support the following HTTP methods for their v2 REST API:
 *  GET, POST, PATCH, DELETE
 */

const (
	defaultRetryMax    = 3
	defaultServiceName = "cloudzero-client-go"
	defaultTimeout     = time.Second * 30
)

var (
	defaultUserAgent = fmt.Sprintf("jjttech/%s/%s (https://github.com/jjttech/%s)", defaultServiceName, version.Version, defaultServiceName)
)

// Client represents a client for communicating with the New Relic APIs.
type Client struct {
	client    *retryablehttp.Client // client represents the underlying HTTP client.
	apiKey    string                // apiKey to send on requests
	userAgent string                // userAgent to send with requests
}

// NewClient is used to create a new instance of Client.
func NewClient(cfg config.Config) (*Client, error) {
	c := http.Client{
		Timeout:   defaultTimeout,
		Transport: http.DefaultTransport,
	}

	if cfg.Timeout != nil {
		c.Timeout = *cfg.Timeout
	}

	if cfg.HTTPTransport != nil {
		c.Transport = cfg.HTTPTransport
	}

	r := retryablehttp.NewClient()
	r.HTTPClient = &c
	r.RetryMax = defaultRetryMax
	r.Logger = nil // Disable logging in go-retryablehttp since we are logging requests directly here

	if cfg.RetryMax != nil {
		r.RetryMax = *cfg.RetryMax
	}

	client := &Client{
		client: r,
	}

	if cfg.UserAgent != "" {
		client.userAgent = cfg.UserAgent
	}

	if cfg.APIKey != "" {
		client.apiKey = cfg.APIKey
	}

	return client, nil
}

// setHeaders applies the default headers needed to talk to the API endpoint
func (c *Client) setHeaders(req *retryablehttp.Request) error {
	if "" != c.apiKey {
		req.Header.Set("Authorization", c.apiKey)
	}

	if "" != c.userAgent {
		req.Header.Set("User-Agent", c.userAgent)
	} else {
		req.Header.Set("User-Agent", defaultUserAgent)
	}

	req.Header.Set("Content-Type", "application/json")

	return nil
}

// Get issues a HTTP Get request against the URL specified
func (c *Client) Get(ctx context.Context, url string) (*http.Response, error) {
	req, err := retryablehttp.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	if err = c.setHeaders(req); err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

// Post issues a HTTP Post request against the URL specified
func (c *Client) Post(ctx context.Context, url string, body interface{}) (*http.Response, error) {
	req, err := retryablehttp.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	if err = c.setHeaders(req); err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

// Put issues a HTTP Put request against the URL specified
func (c *Client) Put(ctx context.Context, url string, body interface{}) (*http.Response, error) {
	req, err := retryablehttp.NewRequestWithContext(ctx, http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}

	if err = c.setHeaders(req); err != nil {
		return nil, err
	}

	return c.client.Do(req)
}

// Delete issues a HTTP Delete request against the URL specified
func (c *Client) Delete(ctx context.Context, url string) (*http.Response, error) {
	req, err := retryablehttp.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	if err = c.setHeaders(req); err != nil {
		return nil, err
	}

	return c.client.Do(req)
}
