package main

import (
	"log"
	"strings"

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
