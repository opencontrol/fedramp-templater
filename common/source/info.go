package source

// Source represents the source of where the information is located.
type Source string

const (
	// SSP indicates that the information is located in the SSP document.
	SSP Source = "SSP"
	// YAML indicates that the information is located in a YAML file.
	YAML Source = "YAML"
)
