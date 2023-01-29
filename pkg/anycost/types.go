package anycost

import (
	"time"
)

const (
	MaxRowsPerFile = 1000000                // MaxRowsPerFile output per docs
	TimeLayout     = "2022-03-16T13:00:00Z" // TimeLayout used for output per docs (ISO8601)
)

type FileDrop struct {
	Root          string     // Root is a folder within the S3 Bucket where your adaptor will write data. This must be a folder and cannot be the root of the S3 Bucket.
	BillingDataID string     // BillingDataID represents an “atom” of billing data which will be added, removed, or updated as a complete unit. Currently it must represent a single month of billing data and be formatted as the first day of a month to the first day of the next month: YYYYMMDD-YYYMMDD. For example, the billing data ID for the month of May 2022 would be: 20220501-20220601.
	DropID        string     // DropID is a unique identifier under which is a complete set of data for this <billing_data_id>. Only one <drop_id> needs to exist. If there is more than one, the “current” drop ID is indicated by the manifest.json. This is useful for versioning. When new data is available for this <billing_data_id> it should be added under a new <drop_id> and the manifest.json updated. The old <drop_id> can then be removed or kept in case it’s necessary to “revert”. To remove all data associated with a <billing_data_id> simply create an empty <drop_id> and point the manifest to that.
	Data          []DataFile // Data contains one or more data files
}

// Manifest represents the manifest.json file used when creating CBF drops
type Manifest struct {
	Version string `json:"version"`
	DropID  string `json:"current_drop_id"`
}

type DataFile struct {
	Filename string
	Rows     []DataRow
}

type DataRow struct {
	LineItem LineItem `csv:"lineitem"`
	Time     Time     `csv:"time"`
	Resource Resource `csv:"resource"`
	Action   Action   `csv:"action"`
	Usage    Usage    `csv:"usage"`
	Cost     Cost     `csv:"cost"`
	Bill     Bill     `csv:"bill"`
}

type LineItemType string

var LineItemTypes = struct {
	Usage                LineItemType
	Tax                  LineItemType
	Support              LineItemType
	Purchase             LineItemType
	CommittedUsePurchase LineItemType
	Discount             LineItemType
	Credit               LineItemType
	Fee                  LineItemType
	Adjustment           LineItemType
}{
	// The most common type and the default. The line item represents a charge for the use of some cloud resource. “Real Cost” includes only Usage line items and uses the first available cost type provided from: amortized_cost, discounted_cost, or cost
	Usage: "Usage",
	// This line item represents any tax charges. Columns in the time, resource, and action categories should only be populated if the tax is associated with applicable Usage changes. Otherwise these columns should be left blank.
	Tax: "Tax",
	// Charges for support or other human services. Columns in the time category should specify the start and end of the support contract.
	Support: "Support",
	// Charge for a one-time purchase or non-usage based subscription; for example, software bought from the AWS Marketplace. The time category columns should represent the span of time over which the purchase applies (for subscriptions with renewal) or the time of the purchase (for one-time charges).
	Purchase: "Purchase",
	// Charges for a "committed use" (RI, savings plan, etc.) purchase. "Committed use" includes any instrument for which payment is made to receive a reduced rate on future usage. This may include one-time upfront purchases or recurring monthly charges. The time category columns should represent the span of time over which the charge applies (for example, one month for monthly recurring charges.) The amortized_cost cost column of this line item should always be zero (if used).
	CommittedUsePurchase: "CommittedUsePurchase",
	// A negative value cost associated with some Usage line item. Columns in the time, resource, and action categories of this line item should match those of the applicable Usage line item and the cost should be negative. The discounted_cost for this line item should be zero if this discount is fully accounted for in the discounted_cost of other Usage line items. (Please note: you must still include discounted_cost if it should be zero)
	Discount: "Discount",
	// A negative value cost not associated with any specific Usage line items. May represent an adjustment due to an error, rounding, refund, etc.
	Credit: "Credit",
	// A positive value charge for which no other line item type applies.
	Fee: "Fee",
	// An alteration made to the bill to correct for some error or rounding issue.
	Adjustment: "Adjustment",
}

type LineItem struct {
	ID          string       `csv:"id"`          // ID uniquely identifies this specific line item in this specific bill.
	Type        LineItemType `csv:"type"`        // Type is the broad category of this line item: Usage, Tax, Discount, etc. Values have special meaning (see below). If not provided, assumed to be Usage.
	Description string       `csv:"description"` // Description specifying additional information about this line item. (optional)
}

type Time struct {
	UsageStart time.Time `csv:"usage_start"` // UsageStart is the hour during which the charged usage applies. This value should be aligned to the start of the hour and will be treated as such. Note, the current version of the Common Billing Format does not currently use usage_end date.
	UsageEnd   time.Time `csv:"usage_end"`   // UsageEnd is the end of a timespan to which the charged usage applies. Note that although this may be specified, currently all changes are treated as occurring in a single hour at usage_start. Future updates to the ingest process may take advantage of the usage_end value.
}

type Tag struct {
	Key   string
	Value string
}

type Resource struct {
	ID          string `csv:"id"`           // ID uniquely identifies the object for which this charge applies. If not provided one will be created with available information.
	Service     string `csv:"service"`      // Service is the category to which this resource belongs. Generally represents different kinds of services provided by the Cloud Provider for which the customer is charged.
	Account     string `csv:"account"`      // Account is the account or project to which this resource belongs (if applicable).
	Region      string `csv:"region"`       // Region to which this resource belongs (if applicable).
	UsageFamily string `csv:"usage_family"` // UsageFamily is commonly a subdivision of resource/service
	Tags        []Tag  `csv:"tag"`          // Tags are extra attributes associated with the resource. This column may appear multiple times with different values
}

type Action struct {
	Operation string `csv:"operation"`  // Operation is the thing that was done to the resource for which you're being charged.
	UsageType string `csv:"usage_type"` // UsageType is commonly a subdivision of action/operation
	Region    string `csv:"region"`     // Region the operation was performed. (May not match the resource/region for cross-region operations.)
	Account   string `csv:"account"`    // Account the operation was performed. (May not match the resource/account for cross-region operations.)
}

type Usage struct {
	Amount float64 `csv:"amount"` // Amount is a numeric value describing an amount consumed/used. E.g. GB stored/transferred, seconds executed, credits consumed
	Units  string  `csv:"units"`  // Units is a description of the units used for usage/amount
}

type Cost struct {
	Cost           float64 `csv:"cost"`            // Cost associated with this line item. (May be negative for line items which represent discounts / credits.)
	DiscountedCost float64 `csv:"discounted_cost"` // DiscountedCost is the net cost associated with this line item after any discounts, credits, or private pricing are applied.
	AmortizedCost  float64 `csv:"amortized_cost"`  // AmortizedCost is the net effective cost associated with this line item after any discounts, credits, or private pricing are applied. It also includes any "committed use" purchases (e.g. AWS RIs or Savings Plans) amortized over all the "Usage" line items to which they apply.
	OnDemandCost   float64 `csv:"on_demand_cost"`  // OnDemandCost is the cost of this line item would have been if no mechanism for reducing costs applied. In other words, it is the cost as if there were no discounts, applicable "committed use" purchases, private pricing, etc. This is often useful for determining one's effective savings rate.
}

type Bill struct {
	InvoiceID string `csv:"invoice_id"` // InvoiceID uniquely identifies a particular bill. A single billing data ID for a single month may include multiple invoices. Ideally this field will not be populated until an Invoice is closed.
}
