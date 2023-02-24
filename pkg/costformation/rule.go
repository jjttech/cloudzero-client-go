package costformation

import (
	"gopkg.in/yaml.v3"
)

// UnmarshalYAML allows for keeping metadata from the yaml file in the struct.
// Currently limited to HeadComment (comment before the object)
func (r *Rule) UnmarshalYAML(value *yaml.Node) error {
	type ruleDecoder Rule // prevent recursion
	var out ruleDecoder

	if err := value.Decode(&out); err != nil {
		return err
	}

	*r = Rule(out)
	r.HeadComment = value.HeadComment

	return nil
}

// MarshalYAML allows for writing metadata from the struct to the yaml file.
// Currently limited to HeadComment (comment before the object)
func (r Rule) MarshalYAML() (interface{}, error) {
	type ruleEncoder Rule // prevent recursion
	var ret yaml.Node

	if err := ret.Encode(ruleEncoder(r)); err != nil {
		return nil, err
	}

	// comment before the object
	ret.HeadComment = r.HeadComment

	return ret, nil
}
