package control

import (
	"fmt"

	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/fedramp-templater/opencontrols"
)

type NarrativeTable struct {
	tbl table
}

func NewNarrativeTable(root xml.Node) NarrativeTable {
	tbl := table{Root: root}
	return NarrativeTable{tbl}
}

func (t *NarrativeTable) SectionRows() ([]xml.Node, error) {
	// skip the header row
	return t.tbl.searchSubtree(`.//w:tr[position() > 1]`)
}

func (t *NarrativeTable) Fill(openControlData opencontrols.Data) (err error) {
	control, err := t.tbl.controlName()
	if err != nil {
		return
	}
	// TODO remove hard-coding
	sectionKey := "b"

	narrative := openControlData.GetNarrative(control, sectionKey)
	fmt.Println(narrative)

	// TODO fill it in

	return
}
