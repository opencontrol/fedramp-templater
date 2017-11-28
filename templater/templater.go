package templater

import (
	"github.com/opencontrol/fedramp-templater/control"
	"github.com/opencontrol/fedramp-templater/opencontrols"
	"github.com/opencontrol/fedramp-templater/reporter"
	"github.com/opencontrol/fedramp-templater/ssp"
	"log"
)

func fillSummaryTables(s *ssp.Document, openControlData opencontrols.Data) error {
	tables, err := s.SummaryTables()
	if err != nil {
		return err
	}
	for _, table := range tables {
		st, err := control.NewSummaryTable(table)
		if err != nil {
			return err
		}
		err = st.Fill(openControlData)
		if err != nil {
			return err
		}
	}

	return nil
}

func fillNarrativeTables(s *ssp.Document, openControlData opencontrols.Data) (err error) {
	tables, err := s.NarrativeTables()
	if err != nil {
		return
	}
	for _, table := range tables {
		ct := control.NewNarrativeTable(table)
		err = ct.Fill(openControlData)
		if err != nil {
			return
		}
	}

	return
}

func fillParameterTables(s *ssp.Document, openControlData opencontrols.Data) (err error) {
	tables, err := s.ParameterTables()
	if err != nil {
		return
	}
	for _, table := range tables {
		ct := control.NewParameterTable(table)
		err = ct.Fill(openControlData)
		if err != nil {
			return
		}
	}

	return
}

// TemplatizeSSP inserts OpenControl data into (i.e. modifies) the provided SSP.
func TemplatizeSSP(s *ssp.Document, openControlData opencontrols.Data) (err error) {
	fillSummaryTables(s, openControlData)
	fillNarrativeTables(s, openControlData)
	fillParameterTables(s, openControlData)
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
		st, err := control.NewSummaryTable(table)
		if err != nil {
			log.Println(err)
			continue
		}
		tableDiffInfo, err := st.Diff(openControlData)
		if err != nil {
			log.Println(err)
			continue
		}
		diffInfo = append(diffInfo, tableDiffInfo...)
	}
	return diffInfo, nil
}
