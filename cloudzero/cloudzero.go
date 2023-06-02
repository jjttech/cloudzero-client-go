package cloudzero

import (
	"github.com/jjttech/cloudzero-client-go/pkg/config"
	"github.com/jjttech/cloudzero-client-go/pkg/costformation"
	"github.com/jjttech/cloudzero-client-go/pkg/telemetry"
)

// New returns an initialized CloudZero client
func New(opts ...ConfigOption) (*CloudZero, error) {
	var err error

	cfg := config.New()

	for _, fn := range opts {
		if nil != fn {
			if err = fn(&cfg); err != nil {
				return nil, err
			}
		}
	}

	client := CloudZero{
		config: cfg,
	}

	if client.CostFormation, err = costformation.New(client.config); err != nil {
		return nil, err
	}

	if client.Telemetry, err = telemetry.New(client.config); err != nil {
		return nil, err
	}

	return &client, nil
}
