package table

import (
	"encoding/json"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nycdavid/boba-party-kit/internal/config"
	"github.com/nycdavid/boba-party-kit/pkg/httpdriver"
)

type (
	JSONData map[string]any

	Model struct {
		data *config.Data
		rows [][]string

		columns []string
		view    *View
	}

	Mod func(*Model)

	SetRowsMsg struct {
		rows [][]string
	}
)

func New(d *config.Data, columns []string) *Model {
	return &Model{
		data:    d,
		columns: columns,
	}
}

func (t *Model) Init() tea.Cmd {
	var rows [][]string
	if t.data.HTTP != "" {
		var jd JSONData
		cl := httpdriver.NewClient()
		res, err := cl.Get(t.data.HTTP, t.data.Auth.Header.BearerEnvVar)
		if err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal(res, &jd); err != nil {
			log.Fatal(err)
		}

		data := jd["data"].(map[string]any)
		budgets := data["budgets"].([]any)
		for _, budget := range budgets {
			b := budget.(map[string]any)
			rowCells := make([]string, len(t.columns))

			for i, col := range t.columns {
				rowCells[i] = b[col].(string)
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
	case SetRowsMsg:
		t.rows = msg.rows
	}

	return t, nil
}

func (t *Model) View() string {
	v := NewView()

	return v.Render(
		t.columns,
		t.rows,
	)
}

func (t *Model) fetchData() {
}
