//go:build unit

package telemetry

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Record", func() {
	Describe("NewRecord", func() {
		It("returns a new record with defaults", func() {
			r := NewRecord()

			Expect(r.Granularity).To(Equal(GranularityTypes.HOURLY))
		})
	})

	Describe("Valid", func() {
		var rec Record

		BeforeEach(func() {
			rec = Record{
				Value:       1,
				Granularity: GranularityTypes.HOURLY,
				Timestamp:   time.Now(),
				ElementName: "test",
				Stream:      "valid-stream-name",
			}
		})

		DescribeTable("record validation tests",
			func(before func(), e error, success bool) {
				before()
				ok, err := rec.Valid()
				Expect(ok).To(Equal(success))

				if success {
					Expect(err).To(BeNil())
				} else {
					Expect(err).To(MatchError(e))
				}
			},
			Entry("valid record - HOURLY", func() {}, nil, true),
			Entry("valid record - DAILY", func() { rec.Granularity = GranularityTypes.DAILY }, nil, true),
			Entry("invalid - Value", func() { rec.Value = 0 }, ErrInvalidValue, false),
			Entry("invalid - Granularity", func() { rec.Granularity = "FOO" }, ErrInvalidGranularity, false),
			Entry("invalid - Timestamp", func() { rec.Timestamp = time.Time{} }, ErrInvalidTimestamp, false),
			Entry("invalid - ElementName", func() { rec.ElementName = "" }, ErrInvalidElementName, false),
			Entry("invalid - Stream", func() { rec.Stream = "" }, ErrInvalidStream, false),
		)
	})
})
