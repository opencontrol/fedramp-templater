package control


// infoSource represents the source of where the information is located.
type infoSource string

const (
	// sspSrc indicates that the information is located in the SSP document.
	sspSrc infoSource = "SSP"
	// yamlSrc indicates that the information is located in a YAML file.
	yamlSrc infoSource = "YAML"
)
