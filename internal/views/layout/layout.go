package layout

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nycdavid/boba-party-kit/pkg/components/searchbar"
)

type (
	Layout struct {
		config     *Config
		components []tea.Model
	}

	Config struct {
		Search *Search `yaml:"search"`
	}

	Init struct {
		Data *Data `yaml:"data"`
	}

	Search struct {
		Init *Init `yaml:"init"`
	}

	Data struct {
		HTTP string `yaml:"http"`
		Auth *Auth  `yaml:"auth"`
	}

	Auth struct {
		Header *Header `yaml:"header"`
	}

	Header struct {
		BearerEnvVar string `yaml:"bearer-env-var"`
	}

	Mod func(*Layout)
)

func New(c *Config) *Layout {
	l := &Layout{
		config:     c,
		components: make([]tea.Model, 0),
	}

	if c.Search != nil {
		l.components = append(l.components, searchbar.New())
	}

	return l
}

func (l *Layout) Init() tea.Cmd {
	return nil
}

func (l *Layout) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return l, tea.Quit
		}
	}

	return l, nil
}

func (l *Layout) View() string {
	return lipgloss.JoinHorizontal(0, l.components[0].View(), "")
}
