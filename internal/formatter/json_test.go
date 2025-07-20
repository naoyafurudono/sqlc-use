package formatter

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/naoyafurudono/sqlc-use/internal/models"
)

func TestJSONFormatter_Format(t *testing.T) {
	tests := []struct {
		name    string
		report  models.UsageReport
		want    string
		wantErr bool
	}{
		{
			name: "single query single operation",
			report: models.UsageReport{
				"GetUser": []models.TableOperation{
					{Operation: "select", Table: "users"},
				},
			},
			want: `{
  "version": "1.0",
  "effects": {
    "GetUser": "{ select[users] }"
  }
}`,
			wantErr: false,
		},
		{
			name: "multiple queries multiple operations",
			report: models.UsageReport{
				"ListOrganizationMember": []models.TableOperation{
					{Operation: "select", Table: "user"},
					{Operation: "select", Table: "member"},
					{Operation: "select", Table: "organization"},
				},
				"AddMember": []models.TableOperation{
					{Operation: "insert", Table: "member"},
				},
			},
			want: `{
  "version": "1.0",
  "effects": {
    "AddMember": "{ insert[member] }",
    "ListOrganizationMember": "{ select[user] | select[member] | select[organization] }"
  }
}`,
			wantErr: false,
		},
		{
			name:    "empty report",
			report:  models.UsageReport{},
			want:    `{
  "version": "1.0",
  "effects": {}
}`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewJSONFormatter()
			got, err := f.Format(tt.report)

			if (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Parse both to normalize JSON formatting
				var gotJSON, wantJSON interface{}
				if err := json.Unmarshal(got, &gotJSON); err != nil {
					t.Fatalf("Failed to parse output JSON: %v", err)
				}
				if err := json.Unmarshal([]byte(tt.want), &wantJSON); err != nil {
					t.Fatalf("Failed to parse expected JSON: %v", err)
				}

				// Re-marshal to ensure consistent formatting
				gotBytes, _ := json.MarshalIndent(gotJSON, "", "  ")
				wantBytes, _ := json.MarshalIndent(wantJSON, "", "  ")

				if !bytes.Equal(gotBytes, wantBytes) {
					t.Errorf("Format() output mismatch\ngot:\n%s\nwant:\n%s", string(gotBytes), string(wantBytes))
				}
			}
		})
	}
}
