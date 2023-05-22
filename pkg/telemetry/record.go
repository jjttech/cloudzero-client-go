package telemetry

func NewRecord() Record {
	return Record{
		Granularity: GranularityTypes.HOURLY,
	}
}

// Valid checks the record to ensure it has all the required fields completed
func (r Record) Valid() (bool, error) {
	if r.Value <= 0 {
		return false, ErrInvalidValue
	}

	if r.Granularity != GranularityTypes.DAILY && r.Granularity != GranularityTypes.HOURLY {
		return false, ErrInvalidGranularity
	}

	if r.Timestamp.IsZero() {
		return false, ErrInvalidTimestamp
	}

	if r.ElementName == "" {
		return false, ErrInvalidElementName
	}

	if !regexpValidStream(r.Stream) {
		return false, ErrInvalidStream
	}

	// Filter isn't a pointer, and {} is valid, so we don't need to validate it here

	return true, nil
}
