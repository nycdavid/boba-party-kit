package layout

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nycdavid/boba-party-kit/internal/config"
	"github.com/nycdavid/boba-party-kit/pkg/components/searchbar"
	"github.com/nycdavid/boba-party-kit/pkg/components/table"
)

type (
	Layout struct {
		config     *config.Config
		components []tea.Model
		focus      int
	}

	Mod func(*Layout)
)

func New(c *config.Config) *Layout {
	l := &Layout{
		config:     c,
		components: make([]tea.Model, 0),
	}

	if c.Search != nil {
		l.components = append(l.components, searchbar.New())

		if c.Search.Results.Table != nil {
			l.components = append(
				l.components,
				table.New(c.Search.Init.Data, c.Search.Results.Table.Columns),
			)
		}
	}

	return l
}

func (l *Layout) Init() tea.Cmd {
	cmds := make([]tea.Cmd, len(l.components))
	for i, c := range l.components {
		cmds[i] = c.Init()
	}

	return tea.Batch(cmds...)
}

func (l *Layout) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return l, tea.Quit
		}
	}

	for _, c := range l.components {
		c.Update(msg)
	}

	return l, nil
}

func (l *Layout) View() string {
	var s string
	for _, comp := range l.components {
		s = lipgloss.JoinVertical(0, s, comp.View())
	}

	return s
}
