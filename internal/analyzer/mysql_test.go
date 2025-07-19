package analyzer

import (
	"reflect"
	"sort"
	"testing"

	"github.com/naoyafurudono/sqlc-use/internal/models"
)

func TestMySQLAnalyzer_Analyze(t *testing.T) {
	tests := []struct {
		name      string
		queryName string
		sql       string
		want      []models.TableOperation
		wantErr   bool
	}{
		{
			name:      "simple select",
			queryName: "GetUser",
			sql:       "SELECT * FROM users WHERE id = ?",
			want: []models.TableOperation{
				{Operation: "select", Table: "users"},
			},
			wantErr: false,
		},
		{
			name:      "select with join",
			queryName: "ListOrganizationMember",
			sql: `SELECT user.* FROM user 
				  INNER JOIN member ON user.id = member.user_id
				  INNER JOIN organization ON organization.id = member.organization_id
				  WHERE organization.name = ?`,
			want: []models.TableOperation{
				{Operation: "select", Table: "user"},
				{Operation: "select", Table: "member"},
				{Operation: "select", Table: "organization"},
			},
			wantErr: false,
		},
		{
			name:      "insert",
			queryName: "AddMember",
			sql:       "INSERT INTO member (user_id, organization_id) VALUES (?, ?)",
			want: []models.TableOperation{
				{Operation: "insert", Table: "member"},
			},
			wantErr: false,
		},
		{
			name:      "update",
			queryName: "UpdateUser",
			sql:       "UPDATE users SET name = ? WHERE id = ?",
			want: []models.TableOperation{
				{Operation: "update", Table: "users"},
			},
			wantErr: false,
		},
		{
			name:      "delete",
			queryName: "RemoveMember",
			sql:       "DELETE FROM member WHERE user_id = ? AND organization_id = ?",
			want: []models.TableOperation{
				{Operation: "delete", Table: "member"},
			},
			wantErr: false,
		},
		{
			name:      "invalid sql",
			queryName: "Invalid",
			sql:       "INVALID SQL STATEMENT",
			want:      nil,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := NewMySQLAnalyzer()
			got, err := analyzer.Analyze(tt.queryName, tt.sql)

			if (err != nil) != tt.wantErr {
				t.Errorf("Analyze() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Sort operations for stable comparison
				sortOperations(got.Operations)
				sortOperations(tt.want)

				if !reflect.DeepEqual(got.Operations, tt.want) {
					t.Errorf("Analyze() operations = %v, want %v", got.Operations, tt.want)
				}
			}
		})
	}
}

func sortOperations(ops []models.TableOperation) {
	sort.Slice(ops, func(i, j int) bool {
		if ops[i].Table != ops[j].Table {
			return ops[i].Table < ops[j].Table
		}
		return ops[i].Operation < ops[j].Operation
	})
}