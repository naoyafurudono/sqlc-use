package formatter

import (
	"encoding/json"

	"github.com/naoyafurudono/sqlc-use/internal/models"
)

// JSONFormatter formats usage reports as JSON
type JSONFormatter struct {
	indent string
}

// NewJSONFormatter creates a new JSON formatter
func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{
		indent: "  ",
	}
}

// Format converts effects report to JSON format
func (f *JSONFormatter) Format(report *models.EffectsReport) ([]byte, error) {
	return json.MarshalIndent(report, "", f.indent)
}
