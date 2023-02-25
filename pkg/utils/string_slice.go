package utils

import (
	"fmt"
	"strconv"

	"gopkg.in/yaml.v3"
)

type StringSlice []string

// Equals does a deep comparison to see if the two String Slices are the same
func (s *StringSlice) Equals(a *StringSlice) bool {
	if nil == a || nil == s {
		return false
	}

	if len(*a) != len(*s) {
		return false
	}

	for i := range *a {
		if (*a)[i] != (*s)[i] {
			return false
		}
	}

	return true
}

// UnmarshalYAML allows for the yaml representation of StringSlice to be either a single
// string or a slice of strings. Does not matter what the actual value type of the item
// contained is (int/bool/etc), it's converted to a string. Uses YAML v3
func (s *StringSlice) UnmarshalYAML(value *yaml.Node) error {
	if nil == value {
		return nil
	}

	var iface interface{}
	if err := value.Decode(&iface); err != nil {
		return err
	}

	var ret StringSlice

	switch str := iface.(type) {
	case bool:
		ret = append(ret, strconv.FormatBool(str))
	case int:
		ret = append(ret, fmt.Sprint(str))
	case float64:
		ret = append(ret, fmt.Sprintf("%.0f", str))
	case string:
		if str != "" {
			ret = append(ret, str)
		}
	case []string:
		ret = make(StringSlice, len(str))
		copy(ret, str)
	case []interface{}:
		ret = make(StringSlice, len(str))
		for k, v := range str {
			ret[k] = fmt.Sprint(v)
		}
	default:
		return fmt.Errorf("unimplemented type %T for StringSlice: %v", str, str)
	}

	*s = ret

	return nil
}

// MarshalYAML turns a StringSlice into either a single yaml key: value or an array. Uses YAML v3
func (s StringSlice) MarshalYAML() (interface{}, error) {
	l := len(s)

	ret := yaml.Node{}

	switch {
	case 1 == l:
		ret.Kind = yaml.ScalarNode
		ret.Style = yamlStyle(s[0])
		ret.Value = s[0]
	case l > 1:
		ret.Kind = yaml.SequenceNode
		ret.Content = make([]*yaml.Node, l)
		for k, v := range s {
			ret.Content[k] = &yaml.Node{
				Kind:  yaml.ScalarNode,
				Style: yamlStyle(v),
				Value: v,
			}
		}
	default:
		// Len must be zero
		return nil, nil
	}

	return ret, nil
}

// yamlStyle looks at a string to identify if it needs quotes, and what type.
func yamlStyle(s string) yaml.Style {
	if _, err := strconv.Atoi(s); err == nil {
		return yaml.SingleQuotedStyle
	}

	return yaml.Style(0)
}
