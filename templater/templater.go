package templater

import (
	"github.com/opencontrol/fedramp-templater/control"
	"github.com/opencontrol/fedramp-templater/ssp"
)

// TemplatizeSsp inserts template tags into (i.e. modifies) the provided Ssp.
func TemplatizeSsp(s *ssp.Document) (err error) {
	tables, err := s.SummaryTables()
	if err != nil {
		return
	}
	for _, table := range tables {
		ct := control.Table{Root: table}
		err = ct.Fill()
		if err != nil {
			return err
		}
	}

	s.UpdateContent()

	return
}
