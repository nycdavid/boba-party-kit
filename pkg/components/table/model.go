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

	SetRowsMsg struct {
		rows [][]string
	}
)

func New(d *config.Init, tableCfg *config.Table) *Model {
	return &Model{
		data:     d,
		tableCfg: tableCfg,
		view:     NewView(),
	}
}

func (t *Model) Init() tea.Cmd {
	var rows [][]string
	if t.data.HTTP.URL != "" {
		cl := httpdriver.NewClient()
		res, err := cl.Get(t.data.HTTP.URL, t.data.HTTP.Auth.Header.BearerEnvVar)
		if err != nil {
			log.Fatal(err)
		}

		var jsonData map[string]any
		if err := json.Unmarshal(res, &jsonData); err != nil {
			log.Fatal(err)
		}

		colFragments := make([]string, len(t.tableCfg.Columns))
		for i, col := range t.tableCfg.Columns {
			colFragments[i] = fmt.Sprintf("%s: %s", col.Name, col.Path)
		}
		joinCols := "{" + strings.Join(colFragments, ", ") + "}"

		q, err := gojq.Parse(t.tableCfg.Rows + " | " + joinCols)
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
			rowCells := make([]string, len(t.tableCfg.Columns))
			for i, col := range t.tableCfg.Columns {
				rowCells[i] = line[col.Name].(string)
			}

			rows = append(rows, rowCells)
		}
	}

	return func() tea.Msg {
		return SetRowsMsg{rows: rows}
	}
}

func (t *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ui.LoseFocusMsg:
		t.view.borderColor = ui.InactiveColor
	case ui.TakeFocusMsg:
		t.view.borderColor = ui.ActiveColor
	case SetRowsMsg:
		t.rows = msg.rows
	}

	return t, nil
}

func (t *Model) View() string {
	cols := make([]string, len(t.tableCfg.Columns))
	for i, col := range t.tableCfg.Columns {
		cols[i] = col.Name
	}

	return t.view.Render(cols, t.rows)
}

func (t *Model) fetchData() {
}
