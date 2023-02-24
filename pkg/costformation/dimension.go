package costformation

import (
	"gopkg.in/yaml.v3"
)

// UnmarshalYAML keeps metadata about the dimension loaded in the struct.
// Currently limited to the HeadComment (comment before the object in yaml)
func (d *Dimension) UnmarshalYAML(value *yaml.Node) error {
	type dimensionDecoder Dimension // prevent recursion
	var out dimensionDecoder

	if err := value.Decode(&out); err != nil {
		return err
	}

	*d = Dimension(out)
	d.HeadComment = value.HeadComment

	return nil
}
