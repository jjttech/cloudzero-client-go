package cloudzero

import (
	"github.com/jjttech/cloudzero-client-go/pkg/config"
	"github.com/jjttech/cloudzero-client-go/pkg/costformation"
	"github.com/jjttech/cloudzero-client-go/pkg/telemetry"
)

const (
	DefaultDefinitionFilename = "definition.yaml"
)

// CloudZero client connection
type CloudZero struct {
	CostFormation *costformation.CostFormation
	Telemetry     *telemetry.Telemetry

	config config.Config
}
