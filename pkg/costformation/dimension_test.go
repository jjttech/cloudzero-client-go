//go:build unit

package costformation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gopkg.in/yaml.v3"
)

var _ = Describe("Dimension", func() {
	Describe("UnmarshalYAML", func() {
		It("should error on invalid data", func() {
			out := Dimension{}
			err := yaml.Unmarshal([]byte("nogood"), &out)
			Expect(err).To(HaveOccurred())
		})

		It("should decode a generic object", func() {
			source := `Name: "test"`
			expected := Dimension{
				Name: "test",
			}
			out := Dimension{}

			err := yaml.Unmarshal([]byte(source), &out)
			Expect(err).To(Succeed())
			Expect(out).To(Equal(expected))
		})

		/*
			It("should decode the head comment", func() {
				source := "# here\nName: \"test\""
				expected := Dimension{
					Name:        "test",
					HeadComment: "# here",
				}
				out := Dimension{}

				err := yaml.Unmarshal([]byte(source), &out)
				Expect(err).To(Succeed())
				Expect(out).To(Equal(expected))
			})
		*/
	})

})
