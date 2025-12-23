package layout

import (
	"log"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nycdavid/boba-party-kit/internal/config"
	"github.com/nycdavid/boba-party-kit/pkg/components/searchbar"
	"github.com/nycdavid/boba-party-kit/pkg/components/table"
	"github.com/nycdavid/boba-party-kit/pkg/components/ui"
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

	if c.Init.NamedSearch != "" {
		l.components = append(l.components, searchbar.New())

		searchName := c.Init.NamedSearch
		i := slices.IndexFunc(c.Searches, func(s *config.Search) bool {
			return s.Name == searchName
		})
		if i == -1 {
			log.Fatalf("invalid named search: %s", c.Init.NamedSearch)
		}

		searchConfig := c.Searches[i]

		if searchConfig.Results.Table != nil {
			l.components = append(
				l.components,
				table.New(searchConfig.Init, searchConfig.Results.Table),
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
	cmp := l.components[l.focus]
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		for _, cmp := range l.components {
			cmp.Update(msg)
		}
	case table.SetTableMsg:
		for _, cmp := range l.components {
			cmp.Update(msg)
		}
	case table.SelectRowMsg:
		row := msg.Row
		l.config.Searches
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			cmp.Update(msg)
		case "j":
			cmp.Update(msg)
		case "tab":
			l.focus = (l.focus + 1) % len(l.components)
			for _, cmp := range l.components {
				cmp.Update(ui.LoseFocusMsg{})
			}
			cmp := l.components[l.focus]
			cmp.Update(ui.TakeFocusMsg{})
		case "ctrl+c", "q":
			return l, tea.Quit
		default:
			cmp := l.components[l.focus]
			cmp.Update(msg)
		}
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
