//go:build unit

package costformation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gopkg.in/yaml.v3"
)

var _ = Describe("Rule", func() {
	Describe("UnmarshalYAML", func() {
		It("should error on invalid data", func() {
			out := Rule{}
			err := yaml.Unmarshal([]byte("nogood"), &out)
			Expect(err).To(HaveOccurred())
		})

		It("should decode a generic object", func() {
			source := `Name: "test"`
			expected := Rule{
				Name: "test",
			}
			out := Rule{}

			err := yaml.Unmarshal([]byte(source), &out)
			Expect(err).To(Succeed())
			Expect(out).To(Equal(expected))
		})

		// Comment handling is not well implemented upstream, so this only appears
		// to work if Rule is a sequence (which is how we use it...)
		It("should decode the head comment", func() {
			source := "# here\n- Name: \"test\""
			expected := []Rule{
				{
					Name:        "test",
					HeadComment: "# here",
				},
			}
			out := []Rule{}

			err := yaml.Unmarshal([]byte(source), &out)
			Expect(err).To(Succeed())
			Expect(out).To(Equal(expected))
		})
	})

})
