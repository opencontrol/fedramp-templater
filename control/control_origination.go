package control

import (
	"github.com/jbowtie/gokogiri/xml"
	"fmt"
)

// Origination prefixes.
const (
	serviceProviderCorporate = "Service Provider Corporate"
	serviceProviderSystemSpecific = "Service Provider System Specific"
	serviceProviderHybrid = "Service Provider Hybrid"
	configuredByCustomer = "Configured by Customer"
	providedByCustomer = "Provided by Customer"
	shared = "Shared"
	inherited = "Inherited"

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
	// Check that we only found the one cell.
	if len(rows) != 1 {
		return nil, fmt.Errorf("Unable to find Control Origination cell")
	}
	// Each checkbox is contained in a paragraph.
	var origins []*checkBox
	paragraphs, err := rows[0].Search(".//w:p")
	for _, paragraph := range paragraphs {
		// 1. Find the box of the checkbox
		checkBox, err := paragraph.Search(".//w:checkBox//w:default")
		if len(checkBox) != 1 || err != nil {
			continue
		}
		// Have to use Attr.
		// Using Attribute does not work when checking the value key.
		// Make sure the length is non zero.
		if len(checkBox[0].Attr(checkBoxAttributeKey)) == 0 {
			continue
		}

		// 2) Find the text next to the checkbox.
		textNodes, err := paragraph.Search(".//w:t")
		if len(textNodes) < 1 || err != nil {
			continue
		}

		// Only construct the checkbox struct if the box and text are found.
		origins = append(origins, newCheckBox(checkBox[0], &textNodes))
	}
	return &controlOrigination{cell: rows[0], origins:origins}, nil
}