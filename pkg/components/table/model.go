package table

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/itchyny/gojq"
	"github.com/nycdavid/boba-party-kit/internal/config"
	"github.com/nycdavid/boba-party-kit/pkg/components/ui"
	"github.com/nycdavid/boba-party-kit/pkg/httpdriver"
)

type (
	Model struct {
		data *config.Init
		rows [][]string

		tableCfg *config.Table
		view     *View
	}

	Mod func(*Model)

	SetTableMsg struct {
		rows    [][]string
		columns []string
	}
)

func New(d *config.Init, tableCfg *config.Table) *Model {
	return &Model{
		data:     d,
		tableCfg: tableCfg,
		view:     NewView(),
	}
}

func (m *Model) Init() tea.Cmd {
	var rows [][]string
	columns := make([]string, len(m.tableCfg.Columns))
	for i, c := range m.tableCfg.Columns {
		columns[i] = c.Name
	}

	if m.data.HTTP.URL != "" {
		cl := httpdriver.NewClient()
		res, err := cl.Get(m.data.HTTP.URL, m.data.HTTP.Auth.Header.BearerEnvVar)
		if err != nil {
			log.Fatal(err)
		}

		var jsonData map[string]any
		if err := json.Unmarshal(res, &jsonData); err != nil {
			log.Fatal(err)
		}

		colFragments := make([]string, len(m.tableCfg.Columns))
		for i, col := range m.tableCfg.Columns {
			colFragments[i] = fmt.Sprintf("%s: %s", col.Name, col.Path)
		}
		joinCols := "{" + strings.Join(colFragments, ", ") + "}"

		q, err := gojq.Parse(m.tableCfg.Rows + " | " + joinCols)
		if err != nil {
			log.Fatal(err)
		}

		var lines []map[string]any
		iter := q.Run(jsonData)
		for {
			v, ok := iter.Next()
			if !ok {
				break
			}

			lines = append(lines, v.(map[string]any))
		}

		for _, line := range lines {
			rowCells := make([]string, len(m.tableCfg.Columns))
			for i, col := range m.tableCfg.Columns {
				rowCells[i] = line[col.Name].(string)
			}

			rows = append(rows, rowCells)
		}
	}

	return func() tea.Msg {
		return SetTableMsg{rows: rows, columns: columns}
	}
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

func (m *Model) fetchData() {
}
