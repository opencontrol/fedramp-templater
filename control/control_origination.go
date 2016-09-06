package control

import (
	"fmt"
	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/fedramp-templater/docx"
	"github.com/opencontrol/fedramp-templater/docx/helper"
	"strings"
)

type controlOrigin uint8

// Origination prefixes.
const (
	noOrigin controlOrigin = iota
	serviceProviderCorporateOrigination
	serviceProviderSystemSpecificOrigination
	serviceProviderHybridOrigination
	configuredByCustomerOrigination
	providedByCustomerOrigination
	sharedOrigination
	inheritedOrigination
)

type originMapping struct {
	yamlMapping string
	docMapping  string
}

func (o originMapping) isDocMappingASubstrOf(value string) bool {
	return strings.Contains(value, o.docMapping)
}

func (o originMapping) isYAMLMappingEqualTo(value string) bool {
	return value == o.yamlMapping
}

func getControlOriginMappings() map[controlOrigin]originMapping {
	return map[controlOrigin]originMapping{
		serviceProviderCorporateOrigination: {
			yamlMapping: "service_provider_corporate",
			docMapping:  "Service Provider Corporate",
		},
		serviceProviderSystemSpecificOrigination: {
			yamlMapping: "service_provided_system_specific",
			docMapping:  "Service Provider System Specific",
		},
		serviceProviderHybridOrigination: {
			yamlMapping: "hybrid",
			docMapping:  "Service Provider Hybrid",
		},
		configuredByCustomerOrigination: {
			yamlMapping: "customer_configured",
			docMapping:  "Configured by Customer",
		},
		providedByCustomerOrigination: {
			yamlMapping: "customer_provided",
			docMapping:  "Provided by Customer",
		},
		sharedOrigination: {
			yamlMapping: "shared",
			docMapping:  "Shared",
		},
		inheritedOrigination: {
			yamlMapping: "inherited",
			docMapping:  "Inherited",
		},
	}
}

type controlOrigination struct {
	cell    xml.Node
	origins map[controlOrigin]*docx.CheckBox
}

func detectControlOriginKeyFromDoc(textNodes []xml.Node) controlOrigin {
	textField := helper.ConcatTextNodes(textNodes)
	controlOriginMappings := getControlOriginMappings()
	for controlOrigin, controlOriginMapping := range controlOriginMappings {
		if controlOriginMapping.isDocMappingASubstrOf(textField) {
			return controlOrigin
		}
	}
	return noOrigin
}

func detectControlOriginKeyFromYAML(text string) controlOrigin {
	controlOriginMappings := getControlOriginMappings()
	for controlOrigin, controlOriginMapping := range controlOriginMappings {
		if controlOriginMapping.isYAMLMappingEqualTo(text) {
			return controlOrigin
		}
	}
	return noOrigin
}

func newControlOrigination(st *SummaryTable) (*controlOrigination, error) {
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
	origins := make(map[controlOrigin]*docx.CheckBox)
	paragraphs, err := rows[0].Search(".//w:p")
	if err != nil {
		return nil, err
	}
	for _, paragraph := range paragraphs {
		// 1. Find the box of the checkbox.
		checkBox, err := docx.FindCheckBoxTag(paragraph)
		if err != nil {
			continue
		}

		// 2. Find the text next to the checkbox.
		textNodes, err := paragraph.Search(".//w:t")
		if len(textNodes) < 1 || err != nil {
			continue
		}

		// 3. Detect the key for the map.
		controlOriginKey := detectControlOriginKeyFromDoc(textNodes)
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
	return &controlOrigination{cell: rows[0], origins: origins}, nil
}
