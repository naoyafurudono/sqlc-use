// Package plugin implements the main sqlc-use plugin logic for analyzing query usage.
package plugin

// Options represents the plugin configuration options from sqlc.yaml
type Options struct {
	// Package specifies the package name to use for query prefixes
	Package string `json:"package"`
	// Format specifies the output format (currently only "json" is supported)
	Format string `json:"format"`
}

// DefaultOptions returns the default plugin options
func DefaultOptions() Options {
	return Options{
		Format: "json",
	}
}
