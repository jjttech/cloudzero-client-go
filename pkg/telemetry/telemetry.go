package telemetry

import (
	"context"
	"net/http"

	czhttp "github.com/jjttech/cloudzero-client-go/internal/http"
	"github.com/jjttech/cloudzero-client-go/pkg/config"
)

const (
	TelemetryPath = "/unit-cost/v1/telemetry"
)

// Telemetry API Client
type Telemetry struct {
	client  *czhttp.Client
	baseURL string
}

// New returns a new instance of the Telemetry API client
func New(cfg config.Config) (*Telemetry, error) {
	client, err := czhttp.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &Telemetry{
		client:  client,
		baseURL: cfg.BaseURL,
	}, nil
}

// DeleteStream issues a API request to delete the specified stream. This should return immediately,
// however the deletion can take some time. During that window a new stream with the same name can not
// be recreated.
func (t *Telemetry) DeleteStream(ctx context.Context, name string) error {
	if nil == t {
		return ErrInvalidTelemetry
	}
	if !regexpValidStream(name) {
		return ErrInvalidStream
	}

	resp, err := t.client.Delete(ctx, t.baseURL+TelemetryPath+"/"+name)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return ErrDeleteFailed
	}

	return nil
}
