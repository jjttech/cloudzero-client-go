package main

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"

	"github.com/jjttech/cloudzero-client-go/internal/version"
)

const (
	AppName        = "CloudZero CLI"
	AppDescription = "CloudZero Command Line Interface"
)

var (
	logLevel string
)

var rootCmd = &cobra.Command{
	Use:               strings.ToLower(AppName),
	Short:             AppName,
	Long:              AppDescription,
	Version:           version.Version,
	DisableAutoGenTag: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Set the log level accordingly
		llv, err := log.ParseLevel(logLevel)
		if err != nil {
			log.WithError(err).Fatal("unsupported log level")
		}
		log.SetLevel(llv)

	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		if err != flag.ErrHelp {
			log.Fatal(err)
		}
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "Logging level: panic, fatal, error, warn, info, debug, or trace")
}
