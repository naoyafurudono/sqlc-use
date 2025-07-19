package models

// TableOperation represents a single table operation in a query
type TableOperation struct {
	Operation string `json:"operation"` // "select", "insert", "update", "delete"
	Table     string `json:"table"`
}

// QueryUsage represents the usage information for a single query
type QueryUsage struct {
	QueryName  string           `json:"-"`
	Operations []TableOperation `json:"operations"`
}

// UsageReport represents the complete usage report for all queries
type UsageReport map[string][]TableOperation
