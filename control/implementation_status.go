package control

import (
	"fmt"
	"github.com/jbowtie/gokogiri/xml"
	"github.com/opencontrol/fedramp-templater/common/status"
	"github.com/opencontrol/fedramp-templater/docx"
	"github.com/opencontrol/fedramp-templater/docx/helper"
	"gopkg.in/fatih/set.v0"
)

type implementationStatus struct {
	cell     xml.Node
	statuses map[status.Key]*docx.CheckBox
}

func (o *implementationStatus) getCheckedStatuses() *set.Set {
	// find the implementation statuses currently checked in the section
	checkedImplementationStatuses := set.New()
	for status, checkbox := range o.statuses {
		if checkbox.IsChecked() {
			checkedImplementationStatuses.Add(status)
		}
	}
	return checkedImplementationStatuses
}

func detectImplementationStatusKeyFromDoc(textNodes []xml.Node) status.Key {
	textField := helper.ConcatTextNodes(textNodes)
	implementationStatusMappings := status.GetSourceMappings()
	for implementationStatus, implementationStatusMapping := range implementationStatusMappings {
		if implementationStatusMapping.IsDocMappingASubstrOf(textField) {
			return implementationStatus
		}
	}
	return status.NoStatus
}

func newImplementationStatus(tbl *table) (*implementationStatus, error) {
	// Find the implementation status row.
	rows, err := tbl.Root.Search(".//w:tc[starts-with(normalize-space(.), 'Implementation Status')]")
	if err != nil {
		return nil, err
	}
	// Check that we only found the one cell.
	if len(rows) != 1 {
		return nil, fmt.Errorf("Unable to find Implementation Status cell")
	}
	// Each checkbox is contained in a paragraph.
	statuses := make(map[status.Key]*docx.CheckBox)
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
		implementationStatusKey := detectImplementationStatusKeyFromDoc(textNodes)
		// if couldn't detect a status, skip.
		if implementationStatusKey == status.NoStatus {
			continue
		}
		// if the status is already in the map, skip.
		_, exists := statuses[implementationStatusKey]
		if exists {
			continue
		}

		// Only construct the checkbox struct if the box and text are found.
		statuses[implementationStatusKey] = docx.NewCheckBox(checkBox, &textNodes)
	}
	return &implementationStatus{cell: rows[0], statuses: statuses}, nil
}
