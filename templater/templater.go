package templater

import (
	"github.com/opencontrol/fedramp-templater/control"
	"github.com/opencontrol/fedramp-templater/opencontrols"
	"github.com/opencontrol/fedramp-templater/reporter"
	"github.com/opencontrol/fedramp-templater/ssp"
)

// TemplatizeSSP inserts OpenControl data into (i.e. modifies) the provided SSP.
func TemplatizeSSP(s *ssp.Document, openControlData opencontrols.Data) (err error) {
	tables, err := s.SummaryTables()
	if err != nil {
		return
	}
	for _, table := range tables {
		ct := control.NewSummaryTable(table)
		err = ct.Fill(openControlData)
		if err != nil {
			return err
		}
	}

	s.UpdateContent()

	return
}

// DiffSSP will find the differences between data in the SSP and the OpenControl data.
func DiffSSP(s *ssp.Document, openControlData opencontrols.Data) ([]reporter.Reporter, error) {
	var diffInfo []reporter.Reporter
	tables, err := s.SummaryTables()
	if err != nil {
		return diffInfo, err
	}
	for _, table := range tables {
		ct := control.NewSummaryTable(table)
		tableDiffInfo, err := ct.Diff(openControlData)
		if err != nil {
			return diffInfo, err
		}
		diffInfo = append(diffInfo, tableDiffInfo...)
	}
	return diffInfo, nil
}
