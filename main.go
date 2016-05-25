package main

import (
  "github.com/opencontrol/doc-template/docx"
)

func main() {
  doc := new(docx.Docx)
  doc.ReadFile("foo.doc")
}
