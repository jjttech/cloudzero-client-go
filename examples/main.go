package main

import (
	"github.com/jjttech/cloudzero-client-go/cloudzero"
	log "github.com/sirupsen/logrus"
)

func main() {
	var (
		cz  *cloudzero.CloudZero
		err error
	)

	if cz, err = cloudzero.New(); err != nil {
		log.WithError(err).Fatal("unable to create CloudZero client")
	}

	// Load from the default filename "definition.yaml" in the current directory
	def, err := cz.CostFormation.Read(cloudzero.DefaultDefinitionFilename)
	if err != nil {
		log.WithError(err).Fatal("unable to load file")
	}

	// Print to the screen
	if err = cz.CostFormation.Write(def, ""); err != nil {
		log.WithError(err).Fatal("unable to write file")
	}
}
