package control

import (
	"github.com/jbowtie/gokogiri/xml"
	"fmt"
	"strings"
	"github.com/opencontrol/fedramp-templater/docx"
	"github.com/opencontrol/fedramp-templater/docx/helper"
)

// Origination prefixes.
const (
	noOrigin = ""
	serviceProviderCorporateOrigination = "Service Provider Corporate"
	serviceProviderSystemSpecificOrigination = "Service Provider System Specific"
	serviceProviderHybridOrigination = "Service Provider Hybrid"
	configuredByCustomerOrigination = "Configured by Customer"
	providedByCustomerOrigination = "Provided by Customer"
	sharedOrigination = "Shared"
	inheritedOrigination = "Inherited"

)

func getAllControlOrigins() []string {
	return []string{
		serviceProviderCorporateOrigination,
		serviceProviderSystemSpecificOrigination,
		serviceProviderHybridOrigination,
		configuredByCustomerOrigination,
		providedByCustomerOrigination,
		sharedOrigination,
		inheritedOrigination,
	}
}

type controlOrigination struct {
	cell xml.Node
	origins map[string]*docx.CheckBox
}

func findControlOriginationBox(paragraph xml.Node) (xml.Node, error) {
	checkBoxes, err := paragraph.Search(".//w:checkBox//w:default")
	if err != nil {
		return nil, err
	} else if len(checkBoxes) != 1 {
		return nil, fmt.Errorf("Unable to find the check box for the control origination.")
	}
	return checkBoxes[0], nil
}

func detectControlOriginKey(textNodes []xml.Node) string {
	textField := helper.ConcatTextNodes(textNodes)
	controlOrigins := getAllControlOrigins()
	for _, controlOrigin := range controlOrigins {
		if strings.Contains(textField, controlOrigin){
			return controlOrigin
		}
	}
	return noOrigin
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
	origins := make(map[string]*docx.CheckBox)
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

		// 3. Detect the key for the map.
		controlOriginKey := detectControlOriginKey(textNodes)
		// if couldn't detect an origin, skip.
		if controlOriginKey == noOrigin {
			continue
		}
		// if the origin is already in the map, skip.
		_, exists := origins[controlOriginKey]
		if exists {
			continue
		}

		// Only construct the checkbox struct if the box and text are found.
		origins[controlOriginKey] = docx.NewCheckBox(checkBox, &textNodes)
	}
	return &controlOrigination{cell: rows[0], origins:origins}, nil
}