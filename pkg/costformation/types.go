package costformation

import (
	"github.com/jjttech/cloudzero-client-go/pkg/utils"
)

// Definition represents the yaml config file uploaded to CloudZero
type Definition struct {
	Dimensions map[string]Dimension `yaml:"Dimensions,omitempty"`

	// Attributes of the file, not content, explicitly ignore on yaml rendering
	Version       string `yaml:"-"` // Version returned via the API
	LastUpdated   string `yaml:"-"` // LastUpdated timestamp return via the API
	LastUpdatedBy string `yaml:"-"` // LastUpdatedBy info returned via the API
	HeadComment   string `yaml:"-"` // HeadComment is inserted at the top of the file
}

// Dimension is a custom configuration for grouping data
type Dimension struct {
	// Name This property provides a user visible name for the custom dimension. This is the name that will be visible in areas such as the Explorer. If not present, the ID will be used as the name.
	Name string `yaml:"Name"`
	// Type This defines the type of custom dimension. Possible values are "Allocation" or "Grouping". If omitted, it will default to "Grouping"
	Type string `yaml:"Type,omitempty"`
	// Hide If set to True, the dimension will not be visible in CloudZero, but the dimension can be used as a source for other dimensions. Otherwise the dimension will be visible to all users. This allows you to create some dimensions that are not meant to be used directly by your users, but rather are meant to be building blocks for other dimensions.
	Hide bool `yaml:"Hide,omitempty"`
	// Disable If set to True this dimension will not be compiled and will not be visible in CloudZero and cannot be used as a source for other dimensions. It allows you to remove the dimension without fully deleting it. If this is not set, it will default to False.
	Disable bool `yaml:"Disable,omitempty"`
	// Child This property identifies the next logical dimension in the hierarchy. Currently, this field's value is used to determine the next GroupBy value when drilling down in the Explorer. Any valid Source is a valid Child. For example, the Child of PaymentOption is Service, meaning when viewing data grouped by PaymentOption the next logical grouping is Service.
	Child string `yaml:"Child,omitempty"`
	// Override Group dimensions only
	Override string `yaml:"Override,omitempty"`
	// DefaultValue This property indicates the name of the Element to place any charge that does not match any rules in the definition.
	DefaultValue string `yaml:"DefaultValue,omitempty"`

	// This is a collection of properties that identify source dimensions as well as ways to manipulate the incoming data before rules and conditions are applied.
	SourceProperties `yaml:",inline,omitempty"`

	// This section defines how the cost allocations will be split. See Adding Allocations for more information on creating allocations.
	AllocateBy        AllocateByStreams `yaml:"AllocateBy,omitempty"` // Deprecated: AllocateBy use AllocateByStreams instead
	AllocateByStreams AllocateByStreams `yaml:"AllocateByStreams,omitempty"`
	AllocateByRules   AllocateByRule    `yaml:"AllocateByRules,omitempty"`

	// Rules This section contains the rules and conditions for grouping the sources into elements.
	Rules []Rule `yaml:"Rules,omitempty"`

	// Metadata
	HeadComment string `yaml:"-"`
}

type SourceProperties struct {
	// CoalesceSources This property determines how multiple sources are handled. If it is set to true, then all specified sources are coalesced and treated as a single source. If set to false, then the sources are treated individually. If it is omitted, then it is set to False.
	CoalesceSources bool `yaml:"CoalesceSources,omitempty"`
	// Source This property identifies one or more source dimensions. It specify a single dimension as a string or multiple sources as a list of strings.
	Source utils.StringSlice `yaml:"Source,omitempty"`
	// Sources This property identifies one or more source dimensions. It specify a single dimension as a string or multiple sources as a list of strings.
	Sources utils.StringSlice `yaml:"Sources,omitempty"`
	// Transforms This property specifies one or more transforms to apply to the sources. Transforms are applied in the order they are specified in this list.
	Transforms []Transform `yaml:"Transforms,omitempty"`
}

// AllocateByStreams An explicit allocation must be composed of a list of telemetry streams
type AllocateByStreams struct {
	// Streams are listed in priority order and it is perfectly fine and expected that stream targets within a dimension will overlap at times. If a particular resource is targeted by multiple streams for the same time period, the higher priority stream will take precedence.
	Streams []string `yaml:"Streams,omitempty"`
}

type AllocateByRule struct {
	AllocationMethod string          `yaml:"AllocationMethod,omitempty"`
	AcrossElements   AcrossElement   `yaml:"AcrossElement,omitempty"`
	SpendToAllocate  SpendToAllocate `yaml:"SpendToAllocate,omitempty"`
}

type AcrossElement struct {
	GroupBy GroupBy `yaml:"GroupBy,omitempty"`
}

type GroupBy struct {
	Source     string      `yaml:"Source"`
	Conditions []Condition `yaml:"Conditions,omitempty"`
}

type SpendToAllocate struct {
	Source     string      `yaml:"Source"`
	Conditions []Condition `yaml:"Conditions,omitempty"`
}

type DimensionRuleType string

var DimensionRuleTypes = struct {
	Group    DimensionRuleType
	GroupBy  DimensionRuleType
	Metadata DimensionRuleType
}{
	Group:    "Group",
	GroupBy:  "GroupBy",
	Metadata: "Metadata",
}

type Rule struct {
	Type   DimensionRuleType `yaml:"Type,omitempty"`
	Name   string            `yaml:"Name,omitempty"`
	Format string            `yaml:"Format,omitempty"`

	SourceProperties `yaml:",inline,omitempty"`

	Conditions []Condition   `yaml:"Conditions,omitempty"`
	Values     []interface{} `yaml:"Values,omitempty"`

	// Metadata
	HeadComment string `yaml:"-"`
}

type Transform struct {
	Type      string `yaml:"Type,omitempty"`
	Delimiter string `yaml:"Delimiter,omitempty"`
	Index     int    `yaml:"Index,omitempty"`
}

type Conditional struct {
	And []Condition `yaml:"And,omitempty"`
	Not []Condition `yaml:"Not,omitempty"`
	Or  []Condition `yaml:"Or,omitempty"`
}

type Condition struct {
	Conditional      `yaml:",inline,omitempty"`
	SourceProperties `yaml:",inline,omitempty"`

	Equals     utils.StringSlice `yaml:"Equals,omitempty"`
	Contains   utils.StringSlice `yaml:"Contains,omitempty"`
	HasValue   utils.StringSlice `yaml:"HasValue,omitempty"`
	BeginsWith utils.StringSlice `yaml:"BeginsWith,omitempty"`
}

/*
 * API Types
 */
type DefinitionVersion struct {
	Version       string `json:"version"`                   // Version unique string identifying this version
	LastUpdated   string `json:"last_updated,omitempty"`    // LastUpdated UTC timestamp for the last update of this entity
	LastUpdatedBy string `json:"last_updated_by,omitempty"` // LastUpdatedBy email, username, or api key that created this entity (optional)
	URI           string `json:"uri,omitempty"`             // Only set when fetching a version
}

type cursor struct {
	Next        string `json:"next_cursor,omitempty"`
	Previous    string `json:"previous_cursor,omitempty"`
	HasNext     bool   `json:"has_next,omitempty"`
	HasPrevious bool   `json:"has_previous,omitempty"`
}

type pagination struct {
	PageCount  int    `json:"page_count,omitempty"`
	ItemCount  int    `json:"item_count"`
	TotalCount int    `json:"total_count"`
	Cursor     cursor `json:"cursor"`
}

type sortOptions struct {
	Keys   []string `json:"sort_keys,omitempty"`
	Orders []string `json:"sort_orders,omitempty"`
}

type sortParam struct {
	Key   string `json:"sort_key,omitempty"`
	Order string `json:"sort_order,omitempty"`
}

type sorting struct {
	Available sortOptions `json:"available,omitempty"`
	Current   []sortParam `json:"current,omitempty"`
}

// defRespVersions is the decoded JSON response body for fetching a list of definition versions
type defRespListVersions struct {
	Pagination pagination          `json:"pagination,omitempty"`
	Sorting    sorting             `json:"sorting,omitempty"`
	Filters    interface{}         `json:"filters,omitempty"`
	TotalCount int                 `json:"total_count,omitempty"`
	Versions   []DefinitionVersion `json:"versions,omitempty"`
}

type defRespGetVersion struct {
	Version DefinitionVersion `json:"version,omitempty"`
}
