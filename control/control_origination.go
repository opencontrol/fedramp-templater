package control

import (
	"fmt"
	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/fedramp-templater/common/origin"
	"github.com/opencontrol/fedramp-templater/docx"
	"github.com/opencontrol/fedramp-templater/docx/helper"
	"gopkg.in/fatih/set.v0"
)

type controlOrigination struct {
	cell    xml.Node
	origins map[origin.Key]*docx.CheckBox
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

func detectControlOriginKeyFromDoc(textNodes []xml.Node) origin.Key {
	textField := helper.ConcatTextNodes(textNodes)
	controlOriginMappings := origin.GetSourceMappings()
	for controlOrigin, controlOriginMapping := range controlOriginMappings {
		if controlOriginMapping.IsDocMappingASubstrOf(textField) {
			return controlOrigin
		}
	}
	return origin.NoOrigin
}

func newControlOrigination(tbl *table) (*controlOrigination, error) {
	// Find the control origination row.
	rows, err := tbl.Root.Search(".//w:tc[starts-with(normalize-space(.), 'Control Origination')]")
	if err != nil {
		return nil, err
	}
	// Check that we only found the one cell.
	if len(rows) != 1 {
		return nil, fmt.Errorf("Unable to find Control Origination cell")
	}
	// Each checkbox is contained in a paragraph.
	origins := make(map[origin.Key]*docx.CheckBox)
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
		if controlOriginKey == origin.NoOrigin {
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
