package telemetry

import (
	"errors"
)

var (
	ErrInvalidTelemetry   = errors.New("invalid Telemetry API client")
	ErrInvalidGranularity = errors.New("invalid Granularity, must be DAILY or HOURLY")
	ErrInvalidValue       = errors.New("invalid Value, must be greater than zero")
	ErrInvalidTimestamp   = errors.New("invalid Timestamp")
	ErrInvalidElementName = errors.New("invalid ElementName")
	ErrInvalidStream      = errors.New("invalid Stream, only alphanumeric characters are allowed along with '_', '.', and '-'")

	ErrDeleteFailed = errors.New("failed to delete telemetry stream")
)
