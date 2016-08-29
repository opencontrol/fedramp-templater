package control

import (
	"github.com/jbowtie/gokogiri/xml"
	"fmt"
)

type controlOrigination struct {
	cell xml.Node
	origins []*checkBox
}

func newControlOrigination(st SummaryTable) (*controlOrigination, error) {
	// Find the control origination row
	rows, err := st.Root.Search(".//w:tc[starts-with(normalize-space(.), 'Control Origination')]")
	if err != nil {
		return nil, err
	}
	if len(rows) != 1 {
		return nil, fmt.Errorf("Unable to find Control Origination cell")
	}
	// Find paragraphs that contain check box
	var origins []*checkBox
	paragraphs, err := rows[0].Search(".//w:p")
	for _, paragraph := range paragraphs {
		checkBox, err := paragraph.Search(".//w:checkBox")
		if len(checkBox) != 1 || err != nil {
			panic(len(checkBox))
			//panic(checkBox[11].Parent().Parent().Parent().Parent().Parent().ToUnformattedXml())
			continue
		}
		textNodes, err := paragraph.Search("//w:t")
		if len(textNodes) < 1 || err != nil {
			panic("hi")
			continue
		}
		origins = append(origins, newCheckBox(checkBox[0], &textNodes))
	}
	panic(len(origins))
	return &controlOrigination{cell: rows[0], origins:origins}, nil
}