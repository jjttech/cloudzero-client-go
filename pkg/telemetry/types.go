package telemetry

import (
	"regexp"
	"time"
)

/*
curl --request POST \
     --url https://api.cloudzero.com/unit-cost/v1/telemetry \
     --header 'Authorization: <cloudzero-api-key>' \
     --header 'content-type: application/json' \
     --data '
{
  "records": [
    {
      "granularity": "HOURLY",    // HOURLY || DAILY
      "value": 12,
      "timestamp": "2020-03-03T15:32:44+00:00",
      "element-name": "Adidas",
      "telemetry-stream": "bytes-consumed-in-x-feature",
      "filter": {
        "k8s_pod": [ "pod1" ],
        "k8s_namespace": [ "my-namespace" ],
        "k8s_cluster": [ "k8s-cluster" ],
        "account": [ "12345" ],
        "service": [ "AWSserviceName" ],
        "usage_family": [ "DataTransfer-Out-Bytes" ],
        "cloud_provider": [ "AWS" ],
        "region": [ "us-west-2" ],

				// These will need custom marshal
        "k8s_label:<labelKey>": [ "label" ],
        "tag:<tagKey>": [ "TagValue" ],
        "custom:<customDimensionName>": [ "DimensionValue" ]
      }
    }
  ]
}
'
*/

// regexpValidStream only alphanumeric characters are allowed along with "_", ".", and "-"
var regexpValidStream = regexp.MustCompile(`^[a-zA-Z\.\_\-]+$`).MatchString

// Granularity is the sample frequency over time.
type Granularity string

var GranularityTypes = struct {
	DAILY  Granularity // DAILY Granularity
	HOURLY Granularity // HOURLY Granularity
}{
	DAILY:  "DAILY",
	HOURLY: "HOURLY",
}

type Record struct {
	Granularity Granularity `json:"granularity"`      // Granularity is the sample frequency over time. Either HOURLY or DAILY
	Value       float64     `json:"value"`            // Value is the number of utilization units consumed. Must be greater than 0.
	Timestamp   time.Time   `json:"timestamp"`        // Timestamp (ISO formatted) of when the usage occurred.
	ElementName string      `json:"element-name"`     // ElementName is used to attribute usage to a specific customer, product, tenant, or other entity. Use either a generated uuid or a human-readable name. This value will become an element of the allocation dimension this telemetry stream is used in.
	Stream      string      `json:"telemetry-stream"` // Stream is a unique name to identify this telemetry stream. Best practice is for this to reflect the metric being used to measure utilization and its units. Only alphanumeric characters are allowed along with "_", ".", and "-". * i.e. gigabyte-seconds-for-x-feature, count_of_records_in_x_database, bytes-consumed-in-x-feature
	Filter      Filter      `json:"filter"`           // Filter is A definition of what portion of your infrastructure this utilization telemetry should be applied to. Note: "filter": "*" or the empty filter set "filter": {} will apply this telemetry to all of your spend.
}

type Filter struct {
	K8sPods        []string `json:"k8s_pod,omitempty"`        //K8sPods is a list of Kubernetes Pod names, as seen in CloudZero.
	K8sNamespaces  []string `json:"k8s_namespace,omitempty"`  // K8sNamespaces is a list of Kubernetes Namespaces, as seen in CloudZero.
	Accounts       []string `json:"account,omitempty"`        // Accounts is list of cloud account IDs for any account.
	Services       []string `json:"service,omitempty"`        // Services is a list of any AWS service or other service in the data for an active billing connection. You can view a list of your current used services on CloudZero
	UsageFamilies  []string `json:"usage_family,omitempty"`   // UsageFamilies is a list of any service charges.
	CloudProviders []string `json:"cloud_provider,omitempty"` // CloudProviders is one or more supported cloud providers
	Regions        []string `json:"region,omitempty"`         // Regions is one or more cloud regions supported by the cloud provider.

	// These will need custom marshal
	K8sLabels        map[string][]string `json:"k8s_label,omitempty"` // K8sLabels is one or more values for a specific Kubernetes Label key. I.e. ”k8s_label:cz-team”: [“research“] would filter to the label “cz-team”: “research”.
	Tags             map[string][]string `json:"tag,omitempty"`       // Tags is one or more values for a specific tag key. I.e. ”tag:cz-environment”: [“prod“] would filter to the tag “cz-environment”: “prod”.
	CustomDimensions map[string][]string `json:"custom,omitempty"`    // CustomDimensions is one or more values for a specific custom dimension key. I.e. custom:environment”: [“prod“] would filter to the custom dimension “environment”: “prod”.
}
