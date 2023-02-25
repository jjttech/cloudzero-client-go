package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/jjttech/cloudzero-client-go/cloudzero"
)

var cmdCostFormation = &cobra.Command{
	Use:     "costformation",
	Short:   "CostFormation related commands",
	Long:    "Work with CostFormation APIs and local files",
	Example: "cloudzero costformation",
}

var formatFilename string
var cmdCostFormationFormat = &cobra.Command{
	Use:     "format",
	Short:   "Format YAML file",
	Long:    "Format the CostFormation definition.yaml file",
	Example: "cloudzero costformation format",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			cz  *cloudzero.CloudZero
			err error
		)

		log.Debug("creating CloudZero client")
		if cz, err = cloudzero.New(); err != nil {
			log.WithError(err).Fatal("unable to create CloudZero client")
		}

		log.WithField("filename", formatFilename).Debug("loading CostFormation definition")
		def, err := cz.CostFormation.Read(formatFilename)
		if err != nil {
			log.WithError(err).WithField("filename", formatFilename).Fatal("unable to load file")
		}

		if err = cz.CostFormation.Write(def, formatFilename); err != nil {
			log.WithError(err).WithField("filename", formatFilename).Fatal("unable to write file")
		}
	},
}

func init() {
	cmdCostFormationFormat.PersistentFlags().StringVarP(&formatFilename, "filename", "f", cloudzero.DefaultDefinitionFilename, "Input file to load")

	cmdCostFormation.AddCommand(cmdCostFormationFormat)

	rootCmd.AddCommand(cmdCostFormation)
}
