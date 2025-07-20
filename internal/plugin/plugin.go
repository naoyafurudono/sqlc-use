// Package plugin implements the main sqlc-use plugin logic for analyzing query usage.
package plugin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/naoyafurudono/sqlc-use/internal/analyzer"
	"github.com/naoyafurudono/sqlc-use/internal/formatter"
	"github.com/naoyafurudono/sqlc-use/internal/models"
	"github.com/sqlc-dev/plugin-sdk-go/plugin"
)

// UsePlugin implements the sqlc plugin interface
type UsePlugin struct {
	analyzerFactory analyzer.Factory
	formatter       formatter.Formatter
}

// New creates a new UsePlugin instance
func New(analyzerFactory analyzer.Factory, formatterImpl formatter.Formatter) *UsePlugin {
	return &UsePlugin{
		analyzerFactory: analyzerFactory,
		formatter:       formatterImpl,
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

	// Parse plugin options
	opts := DefaultOptions()
	if len(req.PluginOptions) > 0 {
		if err := json.Unmarshal(req.PluginOptions, &opts); err != nil {
			return nil, fmt.Errorf("failed to unmarshal plugin options: %w", err)
		}
	}

	// Create analyzer for the engine
	analyzerImpl, err := p.analyzerFactory.Create(engine)
	if err != nil {
		return nil, fmt.Errorf("failed to create analyzer: %w", err)
	}

	// Analyze all queries
	report := models.NewEffectsReport()
	for _, query := range req.Queries {
		effects, analyzeErr := analyzerImpl.Analyze(query.Name, query.Text)
		if analyzeErr != nil {
			return nil, fmt.Errorf("failed to analyze query %s: %w", query.Name, analyzeErr)
		}

		// Use fully qualified name if package is specified
		queryName := query.Name
		if opts.Package != "" {
			queryName = opts.Package + "." + query.Name
		}
		report.Effects[queryName] = effects
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
				Name:     "query-table-operations.json",
				Contents: output,
			},
		},
	}

	return resp, nil
}
