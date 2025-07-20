// Package formatter provides interfaces and implementations for formatting query effects reports.
package formatter

import "github.com/naoyafurudono/sqlc-use/internal/models"

// Formatter defines the interface for output formatters
type Formatter interface {
	// Format converts effects report to output format
	Format(report *models.EffectsReport) ([]byte, error)
}
