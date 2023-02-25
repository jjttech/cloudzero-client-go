package cloudzero

import (
	"github.com/jjttech/cloudzero-client-go/pkg/config"
	"github.com/jjttech/cloudzero-client-go/pkg/costformation"
)

const (
	DefaultDefinitionFilename = "definition.yaml"
)

// CloudZero client connection
type CloudZero struct {
	CostFormation *costformation.CostFormation

	config config.Config
}
