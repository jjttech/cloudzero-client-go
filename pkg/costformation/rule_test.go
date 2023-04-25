//go:build unit

package costformation

import (
	. "github.com/onsi/ginkgo/v2"
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

	Describe("MarshalYAML", func() {
		It("should not error on an empty rule", func() {
			r := Rule{}

			y, err := r.MarshalYAML()
			Expect(err).To(Succeed())
			Expect(y).NotTo(BeNil())
		})

		It("should include a head comment", func() {
			r := Rule{
				HeadComment: "comment",
			}

			a, err := r.MarshalYAML()
			Expect(err).To(Succeed())
			Expect(a).NotTo(BeNil())

			y := a.(yaml.Node)
			Expect(y.HeadComment).To(Equal("comment"))
		})
	})
})
