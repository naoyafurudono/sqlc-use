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
	effectsReport := f.convertToEffectsReport(report)
	return json.MarshalIndent(effectsReport, "", f.indent)
}

// convertToEffectsReport converts the old format to the new effects format
func (f *JSONFormatter) convertToEffectsReport(report models.UsageReport) models.EffectsReport {
	effects := make(map[string]string)
	
	for queryName, operations := range report {
		if len(operations) == 0 {
			effects[queryName] = "{ }"
		} else {
			effectsStr := "{ "
			for i, op := range operations {
				if i > 0 {
					effectsStr += " | "
				}
				effectsStr += op.Operation + "[" + op.Table + "]"
			}
			effectsStr += " }"
			effects[queryName] = effectsStr
		}
	}
	
	return models.EffectsReport{
		Version: "1.0",
		Effects: effects,
	}
}
