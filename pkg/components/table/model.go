package table

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nycdavid/boba-party-kit/internal/config"
	"github.com/nycdavid/boba-party-kit/internal/datadriver"
	"github.com/nycdavid/boba-party-kit/internal/formatdriver"
	"github.com/nycdavid/boba-party-kit/pkg/components/ui"
)

type (
	Model struct {
		data *config.SearchInit

		searchCfg config.Search
		tableCfg  *config.Table

		cfg  *config.Config
		view *View
		name string
	}

	Mod func(*Model)

	SetTableMsg struct {
		rows    [][]string
		columns []string
	}

	SelectRowMsg struct {
		Row []string
	}

	DataDriver interface {
		Fetch() ([]byte, error)
	}

	formatDriver interface {
		Format(data []byte) ([][]string, []string, error)
	}
)

func New(d *config.SearchInit, tableCfg *config.Table, cfg *config.Config, searchCfg config.Search) *Model {
	return &Model{
		searchCfg: searchCfg,
		data:      d,
		tableCfg:  tableCfg,
		cfg:       cfg,
		view:      NewView(),
	}
}

func (m *Model) Init() tea.Cmd {
	var drv DataDriver
	if m.searchCfg.Init.HTTP != nil {
		h := m.searchCfg.Init.HTTP
		drv = datadriver.NewHTTP(h.URL, h.Auth.Header.BearerEnvVar, h.Method)
	} else if m.searchCfg.Init.File != nil {
		f := m.searchCfg.Init.File
		drv = datadriver.NewFile(*f)
	}

	return SetTable(m.searchCfg, drv)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		w := msg.Width - 10
		m.view.tbl.SetWidth(w)
		m.view.resetColumnsWithNewWidth(w)
	case ui.LoseFocusMsg:
		m.view.borderColor = ui.InactiveColor
	case ui.TakeFocusMsg:
		m.view.borderColor = ui.ActiveColor
	case SetTableMsg:
		m.view.setColumns(msg.columns)
		m.view.setRows(msg.rows)
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			i := m.view.tbl.Cursor()
			row := m.view.tbl.Rows()[i]
			return m, m.selectRow(row)
		case "k":
			m.view.tbl.MoveUp(1)
		case "j":
			m.view.tbl.MoveDown(1)
		}
	}

	return m, nil
}

func (m *Model) View() string {
	return m.view.Render()
}

func (m *Model) SetView(v *View) {
	m.view = v
}

func (m *Model) selectRow(row []string) tea.Cmd {
	return func() tea.Msg {
		return SelectRowMsg{Row: row}
	}
}

func SetTable(cfg config.Search, drv DataDriver) tea.Cmd {
	data, err := drv.Fetch()
	if err != nil {
		log.Fatal(err)
	}

	var tblFormatter formatDriver
	if cfg.Results.Table.JSON != nil {
		tc := cfg.Results.Table.JSON
		tblFormatter = formatdriver.NewTableJSON(tc.Rows, tc.Columns)
	} else {
		log.Fatal("no table formatdriver selected")
	}

	rows, columns, err := tblFormatter.Format(data)
	if err != nil {
		log.Fatalf("failed to format table: %v", err)
	}

	return func() tea.Msg {
		return SetTableMsg{rows: rows, columns: columns}
	}
}

func (m *Model) Config() config.Search {
	return m.searchCfg
}
