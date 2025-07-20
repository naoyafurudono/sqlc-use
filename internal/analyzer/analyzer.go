// Package analyzer provides interfaces and implementations for analyzing SQL queries.
package analyzer

// Analyzer defines the interface for SQL query analyzers
type Analyzer interface {
	// Analyze extracts table effects from a SQL query
	Analyze(queryName, sql string) (string, error)
}

// Factory creates analyzers based on the database engine
type Factory interface {
	// Create returns an analyzer for the specified engine
	Create(engine string) (Analyzer, error)
}
