package plugin

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/naoyafurudono/sqlc-use/internal/analyzer"
	"github.com/naoyafurudono/sqlc-use/internal/models"
	"github.com/sqlc-dev/plugin-sdk-go/plugin"
)

// mockAnalyzer is a test double for analyzer
type mockAnalyzer struct {
	analyzeFunc func(queryName, sql string) (*models.QueryUsage, error)
}

func (m *mockAnalyzer) Analyze(queryName, sql string) (*models.QueryUsage, error) {
	if m.analyzeFunc != nil {
		return m.analyzeFunc(queryName, sql)
	}
	return nil, errors.New("not implemented")
}

// mockAnalyzerFactory is a test double for analyzer factory
type mockAnalyzerFactory struct {
	createFunc func(engine string) (analyzer.Analyzer, error)
}

func (m *mockAnalyzerFactory) Create(engine string) (analyzer.Analyzer, error) {
	if m.createFunc != nil {
		return m.createFunc(engine)
	}
	return nil, errors.New("not implemented")
}

// mockFormatter is a test double for formatter
type mockFormatter struct {
	formatFunc func(report models.UsageReport) ([]byte, error)
}

func (m *mockFormatter) Format(report models.UsageReport) ([]byte, error) {
	if m.formatFunc != nil {
		return m.formatFunc(report)
	}
	return nil, errors.New("not implemented")
}

func TestUsePlugin_Generate(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() (*UsePlugin, *plugin.GenerateRequest)
		want    *plugin.GenerateResponse
		wantErr bool
	}{
		{
			name: "successful generation",
			setup: func() (*UsePlugin, *plugin.GenerateRequest) {
				mockAnalyzerImpl := &mockAnalyzer{
					analyzeFunc: func(queryName, sql string) (*models.QueryUsage, error) {
						return &models.QueryUsage{
							QueryName: queryName,
							Operations: []models.TableOperation{
								{Operation: "select", Table: "users"},
							},
						}, nil
					},
				}

				factory := &mockAnalyzerFactory{
					createFunc: func(engine string) (analyzer.Analyzer, error) {
						if engine != "mysql" {
							return nil, errors.New("unsupported engine")
						}
						return mockAnalyzerImpl, nil
					},
				}

				formatter := &mockFormatter{
					formatFunc: func(report models.UsageReport) ([]byte, error) {
						return json.MarshalIndent(report, "", "  ")
					},
				}

				p := New(factory, formatter)
				req := &plugin.GenerateRequest{
					Settings: &plugin.Settings{
						Engine: "mysql",
					},
					Queries: []*plugin.Query{
						{
							Name: "GetUser",
							Text: "SELECT * FROM users WHERE id = ?",
						},
					},
				}

				return p, req
			},
			want: &plugin.GenerateResponse{
				Files: []*plugin.File{
					{
						Name:     "query_usage.json",
						Contents: []byte(`{
  "GetUser": [
    {
      "operation": "select",
      "table": "users"
    }
  ]
}`),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "nil request",
			setup: func() (*UsePlugin, *plugin.GenerateRequest) {
				p := New(nil, nil)
				return p, nil
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "missing engine",
			setup: func() (*UsePlugin, *plugin.GenerateRequest) {
				p := New(nil, nil)
				req := &plugin.GenerateRequest{
					Settings: &plugin.Settings{},
				}
				return p, req
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plugin, req := tt.setup()
			got, err := plugin.Generate(context.Background(), req)

			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if len(got.Files) != len(tt.want.Files) {
					t.Errorf("Generate() got %d files, want %d", len(got.Files), len(tt.want.Files))
					return
				}

				gotJSON := string(got.Files[0].Contents)
				wantJSON := string(tt.want.Files[0].Contents)
				if gotJSON != wantJSON {
					t.Errorf("Generate() output mismatch\ngot:\n%s\nwant:\n%s", gotJSON, wantJSON)
				}
			}
		})
	}
}
