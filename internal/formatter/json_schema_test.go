package formatter

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
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

			// Basic structure validation for new schema format
			// Check version
			version, hasVersion := result["version"]
			if !hasVersion {
				t.Errorf("Output missing 'version' field")
			} else if version != "1.0" {
				t.Errorf("Version is %v, want 1.0", version)
			}

			// Check effects
			effects, hasEffects := result["effects"].(map[string]interface{})
			if !hasEffects {
				t.Errorf("Output missing 'effects' field or it's not an object")
				return
			}

			// Validate each query is in effects
			for queryName, operations := range tt.report {
				effectStr, exists := effects[queryName].(string)
				if !exists {
					t.Errorf("Query %s not found in effects", queryName)
					continue
				}

				// Basic validation that effect string contains expected operations
				for _, op := range operations {
					expectedPattern := op.Operation + "[" + op.Table + "]"
					if !contains(effectStr, expectedPattern) {
						t.Errorf("Effect for %s doesn't contain expected pattern %s",
							queryName, expectedPattern)
					}
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

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
