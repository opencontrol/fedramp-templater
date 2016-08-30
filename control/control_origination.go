package control

import (
	"github.com/jbowtie/gokogiri/xml"
	"fmt"
)

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
		checkBox, err := paragraph.Search(".//w:checkBox//w:default")
		if len(checkBox) != 1 || err != nil {
			continue
		}
		// Have to use Attr. Using Attribute does not work for checking "val"
		if len(checkBox[0].Attr("val")) == 0 {
			continue
		}
		textNodes, err := paragraph.Search(".//w:t")
		if len(textNodes) < 1 || err != nil {
			continue
		}
		origins = append(origins, newCheckBox(checkBox[0], &textNodes))
	}
	return &controlOrigination{cell: rows[0], origins:origins}, nil
}