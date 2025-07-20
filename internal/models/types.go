// Package models defines the data structures used for representing SQL query effects.
package models

// EffectsReport represents the effects report with version and effects
type EffectsReport struct {
	Version string            `json:"version"`
	Effects map[string]string `json:"effects"`
}

// NewEffectsReport creates a new EffectsReport with the current version
func NewEffectsReport() *EffectsReport {
	return &EffectsReport{
		Version: "1.0",
		Effects: make(map[string]string),
	}
}
