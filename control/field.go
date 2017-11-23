package control

import (
	"../common/source"
)

// field is the very simple representation of information.
type field struct {
	source source.Source
	text   string
}
