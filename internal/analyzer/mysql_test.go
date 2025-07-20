package analyzer

import (
	"testing"
)

func TestMySQLAnalyzer_Analyze(t *testing.T) {
	tests := []struct {
		name      string
		queryName string
		sql       string
		want      string
		wantErr   bool
	}{
		{
			name:      "simple select",
			queryName: "GetUser",
			sql:       "SELECT * FROM users WHERE id = ?",
			want: "{ select[users] }",
			wantErr: false,
		},
		{
			name:      "select with join",
			queryName: "ListOrganizationMember",
			sql: `SELECT user.* FROM user
				  INNER JOIN member ON user.id = member.user_id
				  INNER JOIN organization ON organization.id = member.organization_id
				  WHERE organization.name = ?`,
			want: "{ select[member] | select[organization] | select[user] }",
			wantErr: false,
		},
		{
			name:      "insert",
			queryName: "AddMember",
			sql:       "INSERT INTO member (user_id, organization_id) VALUES (?, ?)",
			want: "{ insert[member] }",
			wantErr: false,
		},
		{
			name:      "update",
			queryName: "UpdateUser",
			sql:       "UPDATE users SET name = ? WHERE id = ?",
			want:      "{ update[users] }",
			wantErr: false,
		},
		{
			name:      "delete",
			queryName: "RemoveMember",
			sql:       "DELETE FROM member WHERE user_id = ? AND organization_id = ?",
			want: "{ delete[member] }",
			wantErr: false,
		},
		{
			name:      "invalid sql",
			queryName: "Invalid",
			sql:       "INVALID SQL STATEMENT",
			want:      "",
			wantErr:   true,
		},
		{
			name:      "union simple",
			queryName: "GetActiveAndInactiveUsers",
			sql:       "SELECT * FROM active_users UNION SELECT * FROM inactive_users",
			want: "{ select[active_users] | select[inactive_users] }",
			wantErr: false,
		},
		{
			name:      "union all",
			queryName: "GetAllTransactions",
			sql:       "SELECT * FROM transactions_2023 UNION ALL SELECT * FROM transactions_2024",
			want: "{ select[transactions_2023] | select[transactions_2024] }",
			wantErr: false,
		},
		{
			name:      "union with joins",
			queryName: "GetComplexUnion",
			sql: `SELECT u.id, u.name FROM users u 
				  JOIN orders o ON u.id = o.user_id
				  UNION
				  SELECT c.id, c.name FROM customers c
				  JOIN purchases p ON c.id = p.customer_id`,
			want: "{ select[customers] | select[orders] | select[purchases] | select[users] }",
			wantErr: false,
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
				if got != tt.want {
					t.Errorf("Analyze() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

