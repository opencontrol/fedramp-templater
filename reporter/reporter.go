package reporter

import "io"

// Reporter is construct that reports information in various formats.
//
// WriteTextTo will take the information it has and writes it as plain text to the writer.
type Reporter interface {
	WriteTextTo(io.Writer) error
}
