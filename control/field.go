package control

type field struct {
	source fieldSource
	text   string
}

type fieldSource string

const (
	sspSrc  fieldSource = "SSP"
	yamlSrc fieldSource = "YAML"
)
