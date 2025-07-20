// Package models defines the data structures used for representing SQL query usage information.
package models

// TableOperation represents a single table operation in a query
type TableOperation struct {
	Operation string `json:"operation"` // "select", "insert", "update", "delete"
	Table     string `json:"table"`
}

// QueryTableOp represents the usage information for a single query
type QueryTableOp struct {
	QueryName  string           `json:"-"`
	Operations []TableOperation `json:"operations"`
}

// UsageReport represents the complete usage report for all queries
type UsageReport map[string][]TableOperation

// EffectsReport represents the new schema format with version and effects
type EffectsReport struct {
	Version string            `json:"version"`
	Effects map[string]string `json:"effects"`
}
