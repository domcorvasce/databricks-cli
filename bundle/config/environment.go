package config

// Environment defines overrides for a single environment.
// This structure is recursively merged into the root configuration.
type Environment struct {
	Bundle *Bundle `json:"bundle,omitempty"`

	Workspace *Workspace `json:"workspace,omitempty"`

	Artifacts map[string]*Artifact `json:"artifacts,omitempty"`

	Resources *Resources `json:"resources,omitempty"`
}