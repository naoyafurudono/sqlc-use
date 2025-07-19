package plugin

import (
	"context"
	"fmt"

	"github.com/naoyafurudono/sqlc-use/internal/analyzer"
	"github.com/naoyafurudono/sqlc-use/internal/formatter"
	"github.com/naoyafurudono/sqlc-use/internal/models"
	"github.com/sqlc-dev/plugin-sdk-go/plugin"
)

// UsePlugin implements the sqlc plugin interface
type UsePlugin struct {
	analyzerFactory analyzer.AnalyzerFactory
	formatter       formatter.Formatter
}

// New creates a new UsePlugin instance
func New(analyzerFactory analyzer.AnalyzerFactory, formatter formatter.Formatter) *UsePlugin {
	return &UsePlugin{
		analyzerFactory: analyzerFactory,
		formatter:       formatter,
	}
}

// Generate implements the plugin.Plugin interface
func (p *UsePlugin) Generate(_ context.Context, req *plugin.GenerateRequest) (*plugin.GenerateResponse, error) {
	if req == nil || req.Settings == nil {
		return nil, fmt.Errorf("invalid request: missing settings")
	}

	// Get database engine from settings
	engine := req.Settings.Engine
	if engine == "" {
		return nil, fmt.Errorf("database engine not specified")
	}

	// Create analyzer for the engine
	analyzer, err := p.analyzerFactory.Create(engine)
	if err != nil {
		return nil, fmt.Errorf("failed to create analyzer: %w", err)
	}

	// Analyze all queries
	report := make(models.UsageReport)
	for _, query := range req.Queries {
		usage, analyzeErr := analyzer.Analyze(query.Name, query.Text)
		if analyzeErr != nil {
			return nil, fmt.Errorf("failed to analyze query %s: %w", query.Name, analyzeErr)
		}

		// For now, use query name without package prefix
		// TODO: Parse package name from PluginOptions if needed
		report[query.Name] = usage.Operations
	}

	// Format the report
	output, err := p.formatter.Format(report)
	if err != nil {
		return nil, fmt.Errorf("failed to format output: %w", err)
	}

	// Create response
	resp := &plugin.GenerateResponse{
		Files: []*plugin.File{
			{
				Name:     "query_usage.json",
				Contents: output,
			},
		},
	}

	return resp, nil
}
