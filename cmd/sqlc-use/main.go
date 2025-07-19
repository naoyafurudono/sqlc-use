// Package main is the entry point for the sqlc-use plugin.
package main

import (
	"github.com/naoyafurudono/sqlc-use/internal/analyzer"
	"github.com/naoyafurudono/sqlc-use/internal/formatter"
	"github.com/naoyafurudono/sqlc-use/internal/plugin"
	"github.com/sqlc-dev/plugin-sdk-go/codegen"
)

func main() {
	// Create dependencies
	analyzerFactory := analyzer.NewDefaultFactory()
	jsonFormatter := formatter.NewJSONFormatter()

	// Create plugin
	p := plugin.New(analyzerFactory, jsonFormatter)

	// Run the plugin
	codegen.Run(p.Generate)
}
