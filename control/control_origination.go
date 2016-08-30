package control

import (
	"github.com/jbowtie/gokogiri/xml"
	"fmt"
)

// Origination prefixes.
const (
	serviceProviderCorporateOrigination = "Service Provider Corporate"
	serviceProviderSystemSpecificOrigination = "Service Provider System Specific"
	serviceProviderHybridOrigination = "Service Provider Hybrid"
	configuredByCustomerOrigination = "Configured by Customer"
	providedByCustomerOrigination = "Provided by Customer"
	sharedOrigination = "Shared"
	inheritedOrigination = "Inherited"

)

type controlOrigination struct {
	cell xml.Node
	origins []*checkBox
}

func findControlOriginationBox(paragraph xml.Node) (xml.Node, error) {
	checkBoxes, err := paragraph.Search(".//w:checkBox//w:default")
	if err != nil {
		return nil, err
	} else if len(checkBoxes) != 1 {
		return nil, fmt.Errorf("Unable to find the check box for the control origination.")
	} else if len(checkBoxes[0].Attr(checkBoxAttributeKey)) == 0 {
		// Have to use Attr.
		// Using Attribute does not work when checking the value key.
		// Make sure the length is non zero.
		return nil, fmt.Errorf("Unable to find the check box value attribute.")
	}
	return checkBoxes[0], nil
}

func newControlOrigination(st SummaryTable) (*controlOrigination, error) {
	// Find the control origination row.
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
	if err != nil {
		return nil, err
	}
	for _, paragraph := range paragraphs {
		// 1. Find the box of the checkbox.
		checkBox, err := findControlOriginationBox(paragraph)
		if err != nil {
			continue
		}

		// 2. Find the text next to the checkbox.
		textNodes, err := paragraph.Search(".//w:t")
		if len(textNodes) < 1 || err != nil {
			continue
		}

		// Only construct the checkbox struct if the box and text are found.
		origins = append(origins, newCheckBox(checkBox, &textNodes))
	}
	return &controlOrigination{cell: rows[0], origins:origins}, nil
}