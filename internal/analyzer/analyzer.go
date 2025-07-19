package analyzer

import (
	"github.com/naoyafurudono/sqlc-use/internal/models"
)

// Analyzer defines the interface for SQL query analyzers
type Analyzer interface {
	// Analyze extracts table operations from a SQL query
	Analyze(queryName, sql string) (*models.QueryUsage, error)
}

// AnalyzerFactory creates analyzers based on the database engine
type AnalyzerFactory interface {
	// Create returns an analyzer for the specified engine
	Create(engine string) (Analyzer, error)
}
