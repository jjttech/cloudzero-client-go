//go:build unit

package utils

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"gopkg.in/yaml.v3"
)

var _ = Describe("StringSlice", func() {
	DescribeTable("Equals",
		func(a *StringSlice, b *StringSlice, expected bool) {
			res := a.Equals(b)
			Expect(res).To(Equal(expected))
		},
		Entry("All Nil", nil, nil, false),
		Entry("Nil A", nil, &StringSlice{"test"}, false),
		Entry("Nil B", &StringSlice{"test"}, nil, false),
		Entry("len(a) > len(b)", &StringSlice{"1", "2"}, &StringSlice{"1"}, false),
		Entry("len(b) > len(a)", &StringSlice{"1"}, &StringSlice{"1", "2"}, false),
		Entry("Single element each, not the same", &StringSlice{"1"}, &StringSlice{"2"}, false),
		Entry("Single element each, equal", &StringSlice{"1"}, &StringSlice{"1"}, true),
		Entry("Multiple elements each, not the same", &StringSlice{"1", "3"}, &StringSlice{"2", "4"}, false),
		Entry("Multiple elements each, equal", &StringSlice{"1", "2"}, &StringSlice{"1", "2"}, true),
		Entry("Multiple elements each, same elements, different order", &StringSlice{"1", "2"}, &StringSlice{"2", "1"}, false),
	)

	Describe("yamlStyle", func() {
		It("returns SingleQuotedStyle for numbers", func() {
			style := yamlStyle("100")
			Expect(style).To(Equal(yaml.SingleQuotedStyle))
		})

		It("returns Style(0) (detect) for non-numbers", func() {
			style := yamlStyle("not-a-number")
			Expect(style).To(Equal(yaml.Style(0)))
		})
	})

	Describe("MarshalYAML", func() {
		It("returns nil on empty StringSlice", func() {
			s := StringSlice{}

			a, err := s.MarshalYAML()
			Expect(err).To(Succeed())
			Expect(a).To(BeNil())
		})

		It("encodes a single value to a ScalarNode", func() {
			s := StringSlice{"value"}

			a, err := s.MarshalYAML()
			Expect(err).To(Succeed())
			Expect(a).NotTo(BeNil())

			// Cast
			b := a.(yaml.Node)
			Expect(b.Kind).To(Equal(yaml.ScalarNode))
			Expect(b.Style).To(Equal(yamlStyle("value")))
			Expect(b.Value).To(Equal("value"))
		})

		It("encodes multiple values to a SequenceNode", func() {
			s := StringSlice{"one", "two"}

			a, err := s.MarshalYAML()
			Expect(err).To(Succeed())
			Expect(a).NotTo(BeNil())

			// Cast
			b := a.(yaml.Node)
			Expect(b.Kind).To(Equal(yaml.SequenceNode))
			Expect(len(b.Content)).To(Equal(2))
		})
	})

	Describe("UnmarshalYAML", func() {
		It("returns nil on nil value", func() {
			s := &StringSlice{}

			err := s.UnmarshalYAML(nil)
			Expect(err).To(Succeed())
		})

		It("returns error on bad decode", func() {
			s := &StringSlice{}
			y := yaml.Node{
				Kind: yaml.Kind(^uint32(0)), // Max uint32
			}

			err := s.UnmarshalYAML(&y)
			Expect(err).To(HaveOccurred())
		})
	})

	DescribeTable("UnmarshalYAML different types",
		func(value any, isErr bool) {
			// Create a node to unmarshal
			y := yaml.Node{}
			err := y.Encode(value)
			Expect(err).To(Succeed())

			s := &StringSlice{}

			err = s.UnmarshalYAML(&y)

			if isErr {
				Expect(err).To(HaveOccurred())
			} else {
				Expect(err).To(Succeed())
			}
		},
		Entry("bool", false, false),
		Entry("int", 1, false),
		Entry("float64", float64(1.2), false), // Can not be .0, or yaml.v3 converts it to an int...
		Entry("string", "test", false),
		Entry("string (empty)", "", false),
		Entry("string slice", []string{"one", "two"}, false), // Actually comes across as []interface{} in decoding...
		Entry("yaml Node", yaml.Node{}, true),                // Unsupported conversion
	)
})
