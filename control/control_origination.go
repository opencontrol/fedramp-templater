package control

import (
	"fmt"
	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/fedramp-templater/docx"
	"github.com/opencontrol/fedramp-templater/docx/helper"
	"gopkg.in/fatih/set.v0"
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

type originMapping map[infoSource]string

func (o originMapping) isDocMappingASubstrOf(value string) bool {
	return strings.Contains(value, o[sspSrc])
}

func (o originMapping) isYAMLMappingEqualTo(value string) bool {
	return value == o[yamlSrc]
}

func getControlOriginMappings() map[controlOrigin]originMapping {
	return map[controlOrigin]originMapping{
		serviceProviderCorporateOrigination: {
			yamlSrc: "service_provider_corporate",
			sspSrc:  "Service Provider Corporate",
		},
		serviceProviderSystemSpecificOrigination: {
			yamlSrc: "service_provided_system_specific",
			sspSrc:  "Service Provider System Specific",
		},
		serviceProviderHybridOrigination: {
			yamlSrc: "hybrid",
			sspSrc:  "Service Provider Hybrid",
		},
		configuredByCustomerOrigination: {
			yamlSrc: "customer_configured",
			sspSrc:  "Configured by Customer",
		},
		providedByCustomerOrigination: {
			yamlSrc: "customer_provided",
			sspSrc:  "Provided by Customer",
		},
		sharedOrigination: {
			yamlSrc: "shared",
			sspSrc:  "Shared",
		},
		inheritedOrigination: {
			yamlSrc: "inherited",
			sspSrc:  "Inherited",
		},
	}
}

type controlOrigination struct {
	cell    xml.Node
	origins map[controlOrigin]*docx.CheckBox
}

func (o *controlOrigination) getCheckedOrigins() *set.Set {
	// find the control origins currently checked in the section
	checkedControlOrigins := set.New()
	for origin, checkbox := range o.origins {
		if checkbox.IsChecked() {
			checkedControlOrigins.Add(origin)
		}
	}
	return checkedControlOrigins
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

func getCheckedOriginsFromYAML(yamlControlOriginationData []string) *set.Set {
	// find the control origins currently checked in the section in the YAML.
	yamlControlOrigins := set.New()
	for _, controlOrigin := range yamlControlOriginationData {
		controlOriginKey := detectControlOriginKeyFromYAML(controlOrigin)
		if controlOriginKey == noOrigin {
			continue
		}
		yamlControlOrigins.Add(controlOriginKey)

	}
	return yamlControlOrigins
}
