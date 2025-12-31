package layout

import (
	"log"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nycdavid/boba-party-kit/internal/config"
	"github.com/nycdavid/boba-party-kit/internal/datadriver"
	"github.com/nycdavid/boba-party-kit/pkg/components/searchbar"
	"github.com/nycdavid/boba-party-kit/pkg/components/table"
	"github.com/nycdavid/boba-party-kit/pkg/components/ui"
)

const (
	TypeSearch = "search"
)

type (
	Layout struct {
		config      *config.Config
		components  []tea.Model
		focus       int
		typ         string
		currentData datastore
		stack       []datastore

		searchTable searchTable
	}

	datastore map[string]string

	searchTable interface {
		tea.Model
		Config() config.Search
	}

	Mod func(*Layout)
)

func New(c *config.Config) *Layout {
	l := &Layout{
		config:      c,
		components:  make([]tea.Model, 0),
		currentData: make(map[string]string),
	}

	if c.Init.NamedSearch != "" {
		l.typ = TypeSearch
		l.components = append(l.components, searchbar.New())

		searchName := c.Init.NamedSearch
		i := slices.IndexFunc(c.Searches, func(s config.Search) bool {
			return s.Name == searchName
		})
		if i == -1 {
			log.Fatalf("invalid named search: %s", c.Init.NamedSearch)
		}

		searchConfig := c.Searches[i]

		if searchConfig.Results.Table != nil {
			t := table.New(searchConfig.Init, searchConfig.Results.Table, c, searchConfig)
			l.searchTable = t
			l.components = append(l.components, t)
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
	case tea.WindowSizeMsg, table.SetTableMsg:
		for _, cmp := range l.components {
			cmp.Update(msg)
		}
	case table.SelectRowMsg:
		row := msg.Row
		return l, l.executeSearch(row)
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			_, c := cmp.Update(msg)
			return l, c
		case "j":
			_, c := cmp.Update(msg)
			return l, c
		case "tab":
			l.focus = (l.focus + 1) % len(l.components)
			for _, cmp := range l.components {
				cmp.Update(ui.LoseFocusMsg{})
			}
			cmp := l.components[l.focus]
			_, c := cmp.Update(ui.TakeFocusMsg{})
			return l, c
		case "ctrl+c", "q":
			return l, tea.Quit
		default:
			cmp := l.components[l.focus]
			_, c := cmp.Update(msg)
			return l, c
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

func (l *Layout) executeSearch(row []string) tea.Cmd {
	currentSearchCfg := l.searchTable.Config()
	nextSearchName := currentSearchCfg.Select.NamedSearch

	i := slices.IndexFunc(l.config.Searches, func(s config.Search) bool { return s.Name == nextSearchName })
	if i == -1 {
		log.Fatalf("invalid named search: %s", nextSearchName)
	}

	nextSearchCfg := l.config.Searches[i]

	for k, i := range currentSearchCfg.Select.Datastore {
		l.currentData[k] = row[i]
	}

	args := make([]string, len(nextSearchCfg.Init.Arguments))
	for i, arg := range nextSearchCfg.Init.Arguments {
		args[i] = l.currentData[arg]
	}

	var drv table.DataDriver
	if nextSearchCfg.Init.HTTP != nil {
		h := nextSearchCfg.Init.HTTP
		drv = datadriver.NewHTTP(h.FormattedURL(args), h.Auth.Header.BearerEnvVar, h.Method)
	} else if nextSearchCfg.Init.File != nil {
		f := nextSearchCfg.Init.File
		drv = datadriver.NewFile(*f)
	}

	return table.SetTable(nextSearchCfg, drv)
}
