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

// Format converts usage report to JSON format
func (f *JSONFormatter) Format(report models.UsageReport) ([]byte, error) {
	return json.MarshalIndent(report, "", f.indent)
}
