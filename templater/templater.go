package templater

import (
	"github.com/opencontrol/fedramp-templater/control_table"
	"github.com/opencontrol/fedramp-templater/ssp"
)

// TemplatizeSsp inserts template tags into (i.e. modifies) the provided Ssp.
func TemplatizeSsp(s *ssp.Ssp) (err error) {
	tables, err := s.SummaryTables()
	if err != nil {
		return
	}
	for _, table := range tables {
		ct := control_table.ControlTable{Root: table}
		err = ct.Fill()
		if err != nil {
			return err
		}
	}

	s.UpdateContent()

	return
}
