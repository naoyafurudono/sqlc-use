package analyzer

import (
	"fmt"
	"strings"
)

// DefaultFactory is the default analyzer factory
type DefaultFactory struct{}

// NewDefaultFactory creates a new default factory
func NewDefaultFactory() *DefaultFactory {
	return &DefaultFactory{}
}

// Create returns an analyzer for the specified engine
func (f *DefaultFactory) Create(engine string) (Analyzer, error) {
	switch strings.ToLower(engine) {
	case "mysql":
		return NewMySQLAnalyzer(), nil
	case "postgresql", "postgres":
		return nil, fmt.Errorf("PostgreSQL support not yet implemented")
	case "sqlite", "sqlite3":
		return nil, fmt.Errorf("SQLite support not yet implemented")
	default:
		return nil, fmt.Errorf("unsupported database engine: %s", engine)
	}
}
