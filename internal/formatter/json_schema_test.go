package formatter

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/naoyafurudono/sqlc-use/internal/models"
)

func TestJSONFormatter_OutputMatchesSchema(t *testing.T) {
	// Test that our output format matches the JSON schema
	tests := []struct {
		name   string
		report models.UsageReport
	}{
		{
			name: "simple operations",
			report: models.UsageReport{
				"GetUser": []models.TableOperation{
					{Operation: "select", Table: "users"},
				},
				"CreateUser": []models.TableOperation{
					{Operation: "insert", Table: "users"},
				},
			},
		},
		{
			name: "complex with joins",
			report: models.UsageReport{
				"ListUserPosts": []models.TableOperation{
					{Operation: "select", Table: "users"},
					{Operation: "select", Table: "posts"},
				},
			},
		},
	}

	formatter := NewJSONFormatter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := formatter.Format(tt.report)
			if err != nil {
				t.Fatalf("Format() error = %v", err)
			}

			// Verify it's valid JSON
			var result map[string]interface{}
			if err := json.Unmarshal(output, &result); err != nil {
				t.Errorf("Output is not valid JSON: %v", err)
			}

			// Basic structure validation
			for queryName, operations := range tt.report {
				if _, exists := result[queryName]; !exists {
					t.Errorf("Query %s not found in output", queryName)
				}

				ops, ok := result[queryName].([]interface{})
				if !ok {
					t.Errorf("Query %s operations is not an array", queryName)
					continue
				}

				if len(ops) != len(operations) {
					t.Errorf("Query %s has %d operations in output, want %d", 
						queryName, len(ops), len(operations))
				}
			}
		})
	}
}

func TestJSONSchema_Exists(t *testing.T) {
	// Check that the schema file exists
	schemaPath := filepath.Join("..", "..", "schema", "query-table-operations.schema.json")
	if _, err := os.Stat(schemaPath); os.IsNotExist(err) {
		t.Errorf("JSON schema file not found at %s", schemaPath)
	}
}